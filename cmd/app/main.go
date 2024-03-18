package main

import (
	"film-library-service/configs"
	"film-library-service/internal/handler"
	"film-library-service/internal/repository"
	"film-library-service/internal/usecase"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}

// @title FilmLibraryService
// @version 1.0
// @description API Server for FilmLibraryService

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	db, err := repository.NewPostgresDB(configs.NewConfig())
	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	usecases := usecase.NewUseCase(repos)
	handlers := handler.NewHandler(usecases)

	handlers.RunRoutes(configs.NewConfig().AppPort)
}
