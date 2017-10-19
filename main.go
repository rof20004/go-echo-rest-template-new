package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/rof20004/go-echo-rest-template-new/database"
	"github.com/rof20004/go-echo-rest-template-new/routes"
)

func main() {
	database.InitDatabase()
	database.InitMigrations()
	routes.InitRoutes()
	routes.Server.Logger.Fatal(routes.Server.Start(":8080"))
}
