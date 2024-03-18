package handler

import (
	"encoding/json"
	"film-library-service/internal/entity"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) actorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetActorsHandler(w, r)
	case "POST":
		h.PostActorsHandler(w, r)
	case "PATCH":
		h.UpdateActorsHandler(w, r)
	case "DELETE":
		h.DeleteActorsHandler(w, r)
	}
}

// @Summary Получить список всех актёров
// @Security JWT
// @Tags actors
// @Description Получить список всех актёров
// @ID get-actors
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /actors/ [get]
func (h *Handler) GetActorsHandler(w http.ResponseWriter, r *http.Request) {
	actors, err := h.usecases.Actor.GetAll()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "Произошла ошибка",
		})
		return
	}
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Data: actors,
	})
	return
}

// @Summary Добавить нового актёра
// @Security JWT
// @Tags actors
// @Description Добавить нового актёра
// @ID post-actors
// @Accept  json
// @Produce  json
// @Param input body entity.ActorsInputBody true "actor info"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /actors/ [post]
func (h *Handler) PostActorsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	//
	var actor entity.Actors
	// Парсинг Body
	err := json.NewDecoder(r.Body).Decode(&actor)
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
	if !actor.ValidationGender() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Неверное значение gender_id",
		})
		return
	}
	if !actor.ValidationFirstName() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Слишком коткое имя",
		})
		return
	}
	if !actor.ValidationLastName() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Слишком коткая фамилия",
		})
		return
	}
	// UseCase
	_, err = h.usecases.Actor.Create(actor)
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

// @Summary Обновить информацию об актёре
// @Security JWT
// @Tags actors
// @Description Обновить информацию об актёре
// @ID put-actors
// @Accept  json
// @Produce  json
// @Param actor_id path int true "actor_id"
// @Param lname query string false "Фамилия"
// @Param fname query string false "Имя"
// @Param mname query string false "Отчество"
// @Param gid query int false "ID Гендера"
// @Param dbirth query string false "День рождение | Пример > 2000-01-01"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /actors/{actor_id} [patch]
func (h *Handler) UpdateActorsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	// Получение actorID
	id := r.URL.Path[len("/actors/"):]
	actorID := 0
	fmt.Sscanf(id, "%d", &actorID)
	if actorID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "В запросе не указан actorID",
		})
		return
	}
	// Получение LastName
	lName := r.URL.Query().Get("lname")
	// Получение FirstName
	fName := r.URL.Query().Get("fname")
	// Получение MiddleName
	mName := r.URL.Query().Get("mname")
	// Получение GenderID
	gID := r.URL.Query().Get("gid")
	// Получение DateBirth
	dBirth := r.URL.Query().Get("dbirth")
	//UseCase
	err := h.usecases.Actor.Update(actorID, lName, fName, mName, gID, dBirth)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Не удалось выполнить запрос",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: "ОК",
	})
	return
}

// @Summary Удалить актёра
// @Security JWT
// @Tags actors
// @Description Удалить актёра
// @ID delete-actors
// @Accept  json
// @Produce  json
// @Param actor_id path int true "actor_id"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /actors/{actor_id} [delete]
func (h *Handler) DeleteActorsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка на админа
	if !h.usecases.User.IsUserAdmin(h.userID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Message: "Недостаточно прав",
		})
		return
	}
	// Получение actorID
	id := r.URL.Path[len("/actors/"):]
	actorID := 0
	fmt.Sscanf(id, "%d", &actorID)
	if actorID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "В запросе не указан actorID",
		})
		return
	}
	//UseCase
	err := h.usecases.Actor.Delete(actorID)
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
