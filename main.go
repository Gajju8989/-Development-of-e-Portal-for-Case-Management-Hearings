package main

import (
	"Pc_Build_Service/Internal/database/config"
	"Pc_Build_Service/Internal/database/migrate"
	"Pc_Build_Service/Internal/database/repo"
	"Pc_Build_Service/Internal/router"
	"Pc_Build_Service/Internal/userinterface/handler"
	"Pc_Build_Service/Internal/userinterface/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.InitDB()
	err := migrate.MigrateAll(config.GetDB())
	if err != nil {
		return
	}
	// Initialize MySQL store

	mySqlStore := repo.NewMySqlStore(config.GetDB())
	fmt.Print(mySqlStore)
	// Initialize service
	userService := service.NewUserInterfaceSvc(mySqlStore)

	// Initialize handler
	userHandler := handler.NewUserInterfaceHandler(userService)

	// Initialize router
	myRouter := router.NewRouter(userHandler)

	// Map the routes
	myRouter.MapRoutes()

}
