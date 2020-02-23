package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cuprumbur/weather-service/api"
	"github.com/Cuprumbur/weather-service/configuration"
	"github.com/Cuprumbur/weather-service/docs"
	_ "github.com/Cuprumbur/weather-service/docs"
	"github.com/Cuprumbur/weather-service/storage"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	c := configuration.NewConfig()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", c.Port)
	docs.SwaggerInfo.BasePath = "/"

	db, err := sql.Open("mysql", c.DB.User+":"+c.DB.Pass+"@/"+c.DB.Name)
	if err != nil {
		panic(err.Error())
	}

	s := storage.NewStorage(db)
	a := api.NewAPI(s)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals
		log.Println("performing shutdown...")
		if err := a.Shutdown(); err != nil {
			log.Printf("failed to shutdown sever: %v", err)
		}
	}()

	log.Printf("service is ready to listen on port: %d", c.Port)
	if err := a.Start(c.Port); err != http.ErrServerClosed {
		log.Printf("sever failed: %v", err)
		os.Exit(1)
	}
}

// @title Swagger Example API of the weather server
// @version 0.1
// @description This is a simple weather server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
