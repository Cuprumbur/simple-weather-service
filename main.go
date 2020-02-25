package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Cuprumbur/weather-service/configuration"
	"github.com/Cuprumbur/weather-service/detector/delivery/http"
	detector "github.com/Cuprumbur/weather-service/detector/repository"
	usecase "github.com/Cuprumbur/weather-service/detector/usecase"
	"github.com/Cuprumbur/weather-service/docs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	logEcho "github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
)

// @title Swagger Example API of the weather server
// @version 0.1
// @description This is a simple weather server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	c := configuration.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", c.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	db, err := sql.Open("mysql", c.DB.User+":"+c.DB.Pass+"@/"+c.DB.Name)
	if err != nil {
		panic(err.Error())
	}

	r := detector.NewMySqlDetectorRepository(db)
	u := usecase.NewDetectorUseCase(r)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.SetLevel(logEcho.DEBUG)

	http.SetupDetectorHandler(e, u)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprint(":", c.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
