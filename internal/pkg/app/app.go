package app

import (
	"GoAPI/internal/app/endpoint"
	"GoAPI/internal/storage/postgres"
	"github.com/gin-gonic/gin"
)

type App struct {
	handler   *endpoint.Endpoint
	dbService *postgres.DBService
}

func New() (*App, error) {
	app := &App{}

	app.dbService = postgres.NewDBService()
	app.handler = endpoint.New(app.dbService)
	return app, nil
}

// Старт апи
func (app *App) Start() error {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/GetUsers", app.handler.GetUsers)
	return nil
}
