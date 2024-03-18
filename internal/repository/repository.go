package repository

import (
	"film-library-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetUser(username string, passwordHash string) (entity.User, error)
}

type UserRepo interface {
	IsUserAdmin(userID int) error
}

type ActorRepo interface {
	GetAll() ([]entity.Actors, error)
	GetMoviesByActorID(actorID int) ([]entity.Movies, error)
	Create(actor entity.Actors) (int, error)
	AddMovie(actorID int, movieID int) error
	UpdateLastName(actor entity.Actors) error
	UpdateFirstName(actor entity.Actors) error
	UpdateMiddleName(actor entity.Actors) error
	UpdateGenderID(actor entity.Actors) error
	UpdateDateBirth(actor entity.Actors) error
	DeleteMovies(actorID int) error
	DeleteMovie(actorID int, movieID int) error
	Delete(actorID int) error
}

type MovieRepo interface {
	GetAll(sort string) ([]entity.Movies, error)
	GetByID(movieID int) (entity.Movies, error)
	Search(userQuery string) ([]entity.Movies, error)
	Create(movie entity.Movies) (int, error)
	AddActor(movieID int, actorID int) error
	UpdateTitle(movie entity.Movies) error
	UpdateDescription(movie entity.Movies) error
	UpdateRating(movie entity.Movies) error
	UpdateReleaseDate(movie entity.Movies) error
	Delete(movieID int) error
}

type Repository struct {
	Authorization
	UserRepo
	ActorRepo
	MovieRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPG(db),
		UserRepo:      NewUserPG(db),
		ActorRepo:     NewActorPG(db),
		MovieRepo:     NewMoviePG(db),
	}
}
