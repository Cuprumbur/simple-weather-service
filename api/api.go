package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Cuprumbur/weather-service/storage"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

)

type API struct {
	storage *storage.Storage
	server  *http.Server
}

func NewAPI(storage *storage.Storage) *API {
	return &API{
		storage: storage,
	}
}

func (a *API) Start(port int) error {
	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.bootRouter(),
	}

	return a.server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

func (a *API) bootRouter() *gin.Engine {
	router := gin.Default()

	//detectors
	router.POST("/detectors", a.CreateDetector)
	router.GET("/detectors", a.GetAllDetectors)
	router.GET("/detectors/:id", a.GetDetector)
	router.POST("/detectors/:id", a.UpdateDetector)
	router.DELETE("/detectors/:id", a.DeleteDetector)

	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func write(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("failed to write: %v", err)
	}
}
