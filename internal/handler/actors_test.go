package handler

import (
	"film-library-service/internal/usecase"
	mock_usecase "film-library-service/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetActorsHandler(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockActor)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		args          args
		expStatusCode int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/actors/", nil),
			},
			mockBehavior:  func(s *mock_usecase.MockActor) {},
			expStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			actor := mock_usecase.NewMockActor(c)
			tt.mockBehavior(actor)

			usecases := &usecase.UseCase{Actor: actor}
			handler := NewHandler(usecases)

			http.Handle("/actors/", handler.logging(http.HandlerFunc(handler.actorsHandler)))

			w := httptest.NewRecorder()

			client := &http.Client{}
			client.Do(tt.args.r)
			log.Println(w.Body.String())
			assert.Equal(t, tt.expStatusCode, w.Code)
		})
	}
}
