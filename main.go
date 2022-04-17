package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"

	apiRoutes "go-simple-api/api"
	"go-simple-api/config"
	"go-simple-api/modules/migration"

	welcomeController "go-simple-api/api/v1/welcome"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}
	db := config.NewPostgreConnection()

	appMigration, _ := strconv.ParseBool(os.Getenv("APP_MIGRATION"))
	if appMigration {
		migration.AutoMigration(db)
	}

	welcomeHandler := welcomeController.NewController()

	e := echo.New()

	routes := &apiRoutes.Routes{
		Welcome: welcomeHandler,
	}
	apiRoutes.CreateRoutes(e, routes)

	go func() {
		appConfig := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
		if err := e.Start(appConfig); err != nil {
			log.Println("Shutting Down The REST Server Success")
			os.Exit(0)
		}
	}()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Wait Shutting Down The REST Server ....")
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Fail Shutting Down The REST Server,", err.Error())
	}
}
