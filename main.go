package main

import (
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/db"
	"general_spider_controll_panel/middleware"
	"general_spider_controll_panel/routes"
	"general_spider_controll_panel/utils"
	"github.com/go-co-op/gocron/v2"
	"log"
	"os"
	"time"
)

func main() {
	routesHandler := middleware.Logger(routes.Setup())

	dbHost := utils.Getenv("DB_HOST")
	dbPort := utils.Getenv("DB_PORT")
	dbUser := utils.Getenv("DB_USERNAME")
	dbPass := utils.Getenv("DB_PASSWORD")
	dbName := utils.Getenv("DB_NAME")

	database := db.NewPostgresDB(dbUser, dbPass, dbHost, dbPort, dbName, db.DisableSSL)
	addr := utils.Getenv("SERVER_HOST") + ":" + utils.Getenv("SERVER_PORT")

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	cron, err := gocron.NewScheduler(gocron.WithLocation(time.Local))
	if err != nil {
		logger.Fatal(err)
		return
	}

	server := app.NewApp(addr, routesHandler, database, cron, logger)
	app.Server = server
	logger.Printf("Listening on %s \n", server.Addr)
	server.ListenAndServe()
}
