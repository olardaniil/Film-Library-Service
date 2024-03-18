package handler

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func (h *Handler) authCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Заголовок авторизации отсутствует", http.StatusUnauthorized)
			return
		}
		authorizationParts := strings.Split(authorizationHeader, " ")
		if len(authorizationParts) != 2 || strings.ToLower(authorizationParts[0]) != "bearer" {
			http.Error(w, "Недопустимый заголовок авторизации", http.StatusUnauthorized)
			return
		}
		token := authorizationParts[1]

		userID, err := h.usecases.Authorization.ParseToken(token)
		if err != nil {
			http.Error(w, "Токен авторизации не валиден", http.StatusUnauthorized)
			return
		}
		h.userID = userID
		next.ServeHTTP(w, req)
	})
}
