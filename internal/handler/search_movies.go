package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) searchMoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetSearchMoviesHandler(w, r)
	}
}

// @Summary Поиск фильмов по фрагменту названия фильма или имени актёра
// @Security JWT
// @Tags movies
// @Description Поиск фильмов по фрагменту названия фильма или имени актёра
// @ID get-search-movies
// @Accept  json
// @Produce  json
// @Param q query string True "Запрос на поиск"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /search-movies/ [get]
func (h *Handler) GetSearchMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Get Parameter
	q := r.URL.Query().Get("q")
	log.Println("q:", q)
	// UseCase
	movies, err := h.usecases.Movie.Search(q)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "Не удалось выполнить запрос",
		})
		return
	}
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Data: movies,
	})
	return
}
