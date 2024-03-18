package handler

import (
	"encoding/json"
	"film-library-service/internal/entity"
	"fmt"
	"log"
	"net/http"
)

type LoginInputBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (lb *LoginInputBody) validate() error {
	if lb.Username == "" {
		return fmt.Errorf("Ошибка валидации: поле username не может быть пустым")
	}
	if lb.Password == "" {
		return fmt.Errorf("Ошибка валидации: поле password не может быть пустым")
	}
	return nil
}

// @Summary Авторизация
// @Tags auth
// @Description Авторизация
// @ID post-login
// @Accept  json
// @Produce  json
// @Param input body LoginInputBody true "Тестовые данные: {user:user} без прав администратора и {admin:admin} с правами администратора"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /login [post]
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	// Проверка метода
	if r.Method != "POST" {
		http.Error(w, "Разрешены только POST запросы", http.StatusMethodNotAllowed)
		return
	}
	// Парсинг Body
	var loginBody LoginInputBody
	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Валидация
	err = loginBody.validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// UseCase
	token, err := h.usecases.Authorization.GenerateToken(loginBody.Username, loginBody.Password)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "Пользователь не найден",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Data: entity.User{
			UserName: loginBody.Username,
			Token:    token,
		},
	})
	return
}
