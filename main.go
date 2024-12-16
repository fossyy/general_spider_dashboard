package main

import (
	"context"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/db"
	"general_spider_controll_panel/middleware"
	"general_spider_controll_panel/routes"
	"general_spider_controll_panel/utils"
	"github.com/briandowns/spinner"
	"github.com/go-co-op/gocron/v2"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "Setting up database: "
	s.Start()
	dbHost := utils.Getenv("DB_HOST")
	dbPort := utils.Getenv("DB_PORT")
	dbUser := utils.Getenv("DB_USERNAME")
	dbPass := utils.Getenv("DB_PASSWORD")
	dbName := utils.Getenv("DB_NAME")

	database, err := db.NewPostgresDB(ctx, dbUser, dbPass, dbHost, dbPort, dbName, db.DisableSSL)
	if err != nil {
		s.Stop()
		log.Fatal(err)
		return
	}
	s.Stop()

	s.Prefix = "Setting up logger: "
	s.Start()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s.Stop()

	s.Prefix = "Setting up cron job: "
	s.Start()
	cron, err := gocron.NewScheduler(gocron.WithLocation(time.Local))
	if err != nil {
		s.Stop()
		logger.Fatal(err)
		return
	}
	s.Stop()

	s.Prefix = "Setting up middleware: "
	s.Start()
	routesHandler := middleware.Logger(routes.Setup())
	s.Stop()

	s.Prefix = "Starting up: "
	s.Start()
	addr := utils.Getenv("SERVER_HOST") + ":" + utils.Getenv("SERVER_PORT")
	server := app.NewApp(addr, routesHandler, database, cron, logger)
	app.Server = server
	s.Stop()
	logger.Printf("Listening on %s \n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
		return
	}
}
