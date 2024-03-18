package handler

import (
	"bytes"
	"film-library-service/internal/entity"
	"film-library-service/internal/usecase"
	mock_usecase "film-library-service/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_login(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockAuthorization, user entity.User)

	tests := []struct {
		name          string
		inputBody     string
		inputUser     entity.User
		mockBehavior  mockBehavior
		expStatusCode int
	}{
		// TODO: Add test cases.
		{
			name:      "OK",
			inputBody: `{"user_name": "admin", "password": "admin"}`,
			inputUser: entity.User{
				UserName: "admin",
				Password: "admin",
			},
			mockBehavior: func(s *mock_usecase.MockAuthorization, user entity.User) {
				//s.EXPECT().GenerateToken(user.UserName, user.Password).Return("token", nil)
			},
			expStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_usecase.NewMockAuthorization(c)
			tt.mockBehavior(auth, tt.inputUser)

			usecases := &usecase.UseCase{Authorization: auth}
			handler := NewHandler(usecases)

			http.Handle("/login", handler.logging(http.HandlerFunc(handler.login)))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tt.inputBody))

			client := &http.Client{}
			client.Do(req)

			assert.Equal(t, tt.expStatusCode, w.Code)
		})
	}
}
