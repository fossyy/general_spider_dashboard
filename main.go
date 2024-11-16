package main

import (
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/db"
	"general_spider_controll_panel/routes"
)

func main() {
	server := app.NewApp("localhost:8080", routes.Setup(), db.NewPostgresDB("postgres", "admin", "localhost", "5432", "test", db.DisableSSL))
	fmt.Printf("Listening on %s \n", server.Addr)
	server.ListenAndServe()
}
