package main

import (
	"ginFramework/config"
	"ginFramework/controller"
	_ "ginFramework/docs"
	"ginFramework/helper"
	"ginFramework/model"
	"ginFramework/repository"
	"ginFramework/router"
	"ginFramework/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title Tag Service API
// @version 1.0
// @description A Tag service api in Go usin Gin Framework

// @host localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagRepository := repository.NewTagsREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagRepository, validate)

	// Controller

	tagsController := controller.NewTagsController(tagsService)

	// router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":3333",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
