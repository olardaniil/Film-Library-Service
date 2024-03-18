package usecase

import (
	"film-library-service/internal/entity"
	"film-library-service/internal/repository"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type Authorization interface {
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
	IsUserAdmin(userID int) bool
}

type Actor interface {
	GetAll() ([]entity.Actors, error)
	Create(actor entity.Actors) (int, error)
	Update(actorID int, lName string, fName string, mName string, genderID string, dateBirth string) error
	Delete(actorID int) error
}

type Movie interface {
	GetAll(sort string) ([]entity.Movies, error)
	Search(userQuery string) ([]entity.Movies, error)
	Create(movie entity.Movies) (int, error)
	Update(movieID int, title string, description string, releaseDate string, rating string) error
	Delete(movieID int) error
}

type UseCase struct {
	Authorization
	User
	Actor
	Movie
}

func NewUseCase(repository *repository.Repository) *UseCase {
	return &UseCase{
		Authorization: NewAuthUsecase(repository.Authorization),
		User:          NewUserUsecase(repository.UserRepo),
		Actor:         NewActorUsecase(repository.ActorRepo),
		Movie:         NewMovieUsecase(repository.MovieRepo),
	}
}
