/*
Copyright © 2020 MATACHE MIHAI <matache91mh@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	storeConfig "github.com/curious-kitten/scratch-post/internal/config/store"
	"github.com/curious-kitten/scratch-post/internal/health"
	"github.com/curious-kitten/scratch-post/internal/http/probes"
	"github.com/curious-kitten/scratch-post/internal/http/router"
	"github.com/curious-kitten/scratch-post/internal/info"
	"github.com/curious-kitten/scratch-post/internal/logger"
	"github.com/curious-kitten/scratch-post/internal/store"
	"github.com/curious-kitten/scratch-post/pkg/endpoints"
	"github.com/curious-kitten/scratch-post/pkg/metadata"
	"github.com/curious-kitten/scratch-post/pkg/projects"
	"github.com/curious-kitten/scratch-post/pkg/scenarios"
)

var (
	port     *string
	dbconfig *string
)

func init() {
	port = flag.String("port", "9090", "Port of server")
	dbconfig = flag.String("dbconfig", "/etc/db.json", "Path to DB config settings")
	flag.Parse()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := info.App{
		Name:    "scratch-post",
		Version: "0.0.1-beta",
	}

	instance := info.InstanceInfo()

	log, flush, err := logger.New(app, instance, true)
	if err != nil {
		panic(err)
	}
	defer flush()

	log.Info("Reading configurations...")
	confContents, err := os.Open(*dbconfig)
	if err != nil {
		panic(err)
	}

	config, err := storeConfig.NewConfig(confContents)
	if err != nil {
		panic(err)
	}
	log.Info("Starting app...")
	r := router.New(log)
	versionedRouter := r.PathPrefix("/api/v1").Subrouter()

	conditions := health.NewConditions(app, instance)
	probes.RegisterHTTPProbes(versionedRouter.PathPrefix("/probes").Subrouter(), conditions)

	meta := metadata.NewMetaManager()

	client, err := store.Client(ctx, config.Address)
	if err != nil {
		panic(err)
	}
	//  Projects endpoint
	projectsCollection, err := store.Collection(config.DataBase, config.Collections.Projects, client)
	if err != nil {
		panic(err)
	}
	projectRouter := versionedRouter.PathPrefix("/projects").Subrouter()
	endpoints.Creator(ctx, projects.New(meta, projectsCollection), projectRouter)
	endpoints.Lister(ctx, projects.List(projectsCollection), projectRouter)
	endpoints.Getter(ctx, projects.Get(projectsCollection), projectRouter)
	endpoints.Deleter(ctx, projects.Delete(projectsCollection), projectRouter)
	endpoints.Updater(ctx, projects.Update(projectsCollection), projectRouter)

	// Scenario endpoints
	scenarioCollection, err := store.Collection(config.DataBase, config.Collections.Scenarios, client)
	if err != nil {
		panic(err)
	}
	scenarioRouter := versionedRouter.PathPrefix("/scenarios").Subrouter()
	endpoints.Creator(ctx, scenarios.New(meta, scenarioCollection, projects.Get(projectsCollection)), scenarioRouter)
	endpoints.Lister(ctx, scenarios.List(scenarioCollection), scenarioRouter)
	endpoints.Getter(ctx, scenarios.Get(scenarioCollection), scenarioRouter)
	endpoints.Deleter(ctx, scenarios.Delete(scenarioCollection), scenarioRouter)
	endpoints.Updater(ctx, scenarios.Update(scenarioCollection, projects.Get(projectsCollection)), scenarioRouter)

	// Start HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Infof("Server started on port %s", *port)

	<-c

	log.Info("Shutting down...")
	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
