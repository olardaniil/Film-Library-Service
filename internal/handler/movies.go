package handler

import (
	"encoding/json"
	"film-library-service/internal/entity"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) moviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetMoviesHandler(w, r)
	case "POST":
		h.PostMoviesHandler(w, r)
	case "PATCH":
		h.UpdateMoviesHandler(w, r)
	case "DELETE":
		h.DeleteMoviesHandler(w, r)
	}
}

// @Summary Получить список всех фильмов
// @Security JWT
// @Tags movies
// @Description Получить список всех актёров
// @ID get-movies
// @Accept  json
// @Produce  json
// @Param sort query string false "Сортировка | Возможные значения: title-asc, title-desc, release-date-asc, release-date-desc, rating-asc, rating-desc"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies/ [get]
func (h *Handler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Get Parameter
	sortingParameter := r.URL.Query().Get("sort")
	// UseCase
	movies, err := h.usecases.Movie.GetAll(sortingParameter)
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

// @Summary Добавить новый фильм
// @Security JWT
// @Tags movies
// @Description Добавить новый фильм
// @ID post-movies
// @Accept  json
// @Produce  json
// @Param input body entity.MoviesInputBody true "movies body"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies/ [post]
func (h *Handler) PostMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	// Парсинг Body
	var bodyMovie entity.Movies
	err := json.NewDecoder(r.Body).Decode(&bodyMovie)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Неверный формат body",
		})
		return
	}
	// Валидация
	err = bodyMovie.Validation()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: err.Error(),
		})
		return
	}
	// UseCase
	_, err = h.usecases.Movie.Create(bodyMovie)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Не удалось выполнить запрос",
		})
		return
	}
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: "ОК",
	})
	return
}

// @Summary Обновить информацию о фильме
// @Security JWT
// @Tags movies
// @Description Обновить информацию о фильме
// @ID put-movies
// @Accept  json
// @Produce  json
// @Param movie_id path int true "movie_id"
// @Param title query string false "Название"
// @Param description query string false "Описание"
// @Param rdate query string false "Дата выхода | Пример > 2000-01-01"
// @Param rating query number false "Рейтинг"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies/{movie_id} [patch]
func (h *Handler) UpdateMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	// Получение movieID
	id := r.URL.Path[len("/movies/"):]
	movieID := 0
	fmt.Sscanf(id, "%d", &movieID)
	if movieID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "В запросе не указан movieID",
		})
		return
	}
	// Получение Title
	title := r.URL.Query().Get("title")
	// Получение Description
	description := r.URL.Query().Get("description")
	// Получение ReleaseDate
	releaseDate := r.URL.Query().Get("rdate")
	// Получение Rating
	rating := r.URL.Query().Get("rating")
	//UseCase
	err := h.usecases.Movie.Update(movieID, title, description, releaseDate, rating)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Не удалось выполнить запрос",
		})
		return
	}
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: "ОК",
	})
	return
}

// @Summary Удалить фильм
// @Security JWT
// @Tags movies
// @Description Удалить фильм
// @ID delete-movies
// @Accept  json
// @Produce  json
// @Param movie_id path int true "movie_id"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies/{movie_id} [delete]
func (h *Handler) DeleteMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	// Получение movieID
	id := r.URL.Path[len("/movies/"):]
	movieID := 0
	fmt.Sscanf(id, "%d", &movieID)
	if movieID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "В запросе не указан movieID",
		})
		return
	}
	//UseCase
	err := h.usecases.Movie.Delete(movieID)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Не удалось выполнить запрос",
		})
		return
	}
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: "ОК",
	})
	return
}
