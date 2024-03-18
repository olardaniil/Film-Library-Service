package repository

import (
	"film-library-service/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ActorPG struct {
	db *sqlx.DB
}

// NewActorPG
func NewActorPG(db *sqlx.DB) *ActorPG {
	return &ActorPG{db: db}
}

// GetAll
func (r *ActorPG) GetAll() ([]entity.Actors, error) {
	query := fmt.Sprintf("SELECT * FROM actors")
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	//
	defer rows.Close()
	var actors []entity.Actors
	for rows.Next() {
		var actor entity.Actors
		err := rows.Scan(&actor.ID, &actor.LastName, &actor.FirstName, &actor.MiddleName, &actor.GenderID, &actor.DateBirth)
		if err != nil {
			return nil, err
		}
		actor.SetGenderName()
		actors = append(actors, actor)
	}
	//
	if err := rows.Err(); err != nil {
		return nil, err
	}
	//
	return actors, nil
}

// GetMoviesByActorID
func (r *ActorPG) GetMoviesByActorID(actorID int) ([]entity.Movies, error) {
	//
	query := fmt.Sprintf("SELECT m.ID, m.title, m.description, m.rating, m.release_date  FROM movies m JOIN actors_movies am on m.id = am.movie_id WHERE am.actor_id=$1")
	rows, err := r.db.Query(query, actorID)
	if err != nil {
		return nil, err
	}
	//
	defer rows.Close()
	var movies []entity.Movies
	for rows.Next() {
		var movie entity.Movies
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	//
	if err := rows.Err(); err != nil {
		return nil, err
	}
	//
	return movies, nil
}

// Create
func (r *ActorPG) Create(actor entity.Actors) (int, error) {
	//
	var id int
	query := fmt.Sprintf("INSERT INTO actors (last_name, first_name, middle_name, gender_id, date_birth) values ($1, $2, $3, $4, $5) RETURNING id")
	//
	row := r.db.QueryRow(query, actor.LastName, actor.FirstName, actor.MiddleName, actor.GenderID, actor.DateBirth)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	//
	return id, nil
}

// AddMovie
func (r *ActorPG) AddMovie(actorID int, movieID int) error {
	query := fmt.Sprintf("INSERT INTO actors_movies (actor_id, movie_id) values ($1, $2)")
	err := r.db.QueryRow(query, actorID, movieID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *ActorPG) UpdateLastName(actor entity.Actors) error {
	query := fmt.Sprintf("UPDATE actors SET last_name=$1 WHERE id=$2")
	_, err := r.db.Exec(query, actor.LastName, actor.ID)
	return err
}

func (r *ActorPG) UpdateFirstName(actor entity.Actors) error {
	query := fmt.Sprintf("UPDATE actors SET first_name=$1 WHERE id=$2")
	_, err := r.db.Exec(query, actor.FirstName, actor.ID)
	return err
}

func (r *ActorPG) UpdateMiddleName(actor entity.Actors) error {
	query := fmt.Sprintf("UPDATE actors SET middle_name=$1 WHERE id=$2")
	_, err := r.db.Exec(query, actor.MiddleName, actor.ID)
	return err
}

func (r *ActorPG) UpdateGenderID(actor entity.Actors) error {
	query := fmt.Sprintf("UPDATE actors SET gender_id=$1 WHERE id=$2")
	_, err := r.db.Exec(query, actor.GenderID, actor.ID)
	return err
}

func (r *ActorPG) UpdateDateBirth(actor entity.Actors) error {
	query := fmt.Sprintf("UPDATE actors SET date_birth=$1 WHERE id=$2")
	_, err := r.db.Exec(query, actor.DateBirth, actor.ID)
	return err
}

func (r *ActorPG) DeleteMovies(actorID int) error {
	query := fmt.Sprintf("DELETE FROM actors_movies WHERE actor_id=$1")
	_, err := r.db.Exec(query, actorID)
	return err
}

func (r *ActorPG) DeleteMovie(actorID int, movieID int) error {
	query := fmt.Sprintf("DELETE FROM actors_movies WHERE actor_id=$1 AND movie_id=$2")
	_, err := r.db.Exec(query, actorID, movieID)
	return err
}

func (r *ActorPG) Delete(actorID int) error {
	err := r.DeleteMovies(actorID)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM actors WHERE id=$1")
	_, err = r.db.Exec(query, actorID)
	return err
}
