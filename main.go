package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cuprumbur/weather-service/api"
	"github.com/Cuprumbur/weather-service/configuration"
	"github.com/Cuprumbur/weather-service/storage"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	c := configuration.NewConfig()

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
