package handler

import (
	_ "film-library-service/docs"
	"film-library-service/internal/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Handler struct {
	usecases *usecase.UseCase
	userID   int
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) RunRoutes(port string) {
	// login
	http.Handle("/login", h.logging(http.HandlerFunc(h.login)))
	// actors
	http.Handle("/actors/", h.logging(h.authCheck(http.HandlerFunc(h.actorsHandler))))
	// movies
	http.Handle("/movies/", h.logging(h.authCheck(http.HandlerFunc(h.moviesHandler))))
	// search-movies
	http.Handle("/search-movies/", h.logging(h.authCheck(http.HandlerFunc(h.searchMoviesHandler))))
	// swagger
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // Путь к файлу документации Swagger JSON
	))
	//
	log.Println("Swagger UI доступен по адресу http://localhost:8080/swagger/index.html")
	//
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
