package main

import (
	"be9/restclean/config"
	"be9/restclean/factory"
	"be9/restclean/migration"
	"be9/restclean/routes"
)

func main() {
	//initiate db connection
	dbConn := config.InitDB()

	// run auto migrate from gorm
	migration.InitMigrate(dbConn)

	// initiate factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":80"))
}
