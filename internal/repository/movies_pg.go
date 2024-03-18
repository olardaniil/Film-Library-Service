package repository

import (
	"film-library-service/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type MoviePG struct {
	db *sqlx.DB
}

func NewMoviePG(db *sqlx.DB) *MoviePG {
	return &MoviePG{db: db}
}

func (r *MoviePG) GetAll(sort string) ([]entity.Movies, error) {
	var query string
	switch sort {

	case "title-asc":
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY title ASC")
	case "title-desc":
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY title DESC")
	case "release-date-asc":
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY release_date ASC")
	case "release-date-desc":
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY release_date DESC")
	case "rating-asc":
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY rating ASC")
	default:
		query = fmt.Sprintf("SELECT * FROM movies ORDER BY rating DESC")
	}

	rows, err := r.db.Query(query)
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

func (r *MoviePG) GetByID(movieID int) (entity.Movies, error) {
	var movie entity.Movies
	query := fmt.Sprintf("SELECT * FROM movies WHERE id=$1")
	err := r.db.Get(&movie, query, movieID)

	return movie, err
}

func (r *MoviePG) Search(userQuery string) ([]entity.Movies, error) {
	query := fmt.Sprintf("SELECT m.id, m.title, m.description, m.release_date, m.rating FROM movies m LEFT JOIN actors_movies am ON m.id = am.movie_id LEFT JOIN actors a ON am.actor_id = a.id WHERE lower(m.title) LIKE $1 OR lower(a.first_name) LIKE $1 OR lower(a.last_name) LIKE $1 GROUP BY m.id")
	rows, err := r.db.Query(query, "%"+strings.ToLower(userQuery)+"%")
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//
	var movies []entity.Movies
	for rows.Next() {
		log.Println(1)
		var movie entity.Movies
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func (r *MoviePG) Create(movie entity.Movies) (int, error) {
	//
	var id int
	query := fmt.Sprintf("INSERT INTO movies (title, description, rating, release_date) values ($1, $2, $3, $4) RETURNING id")
	//
	row := r.db.QueryRow(query, movie.Title, movie.Description, movie.Rating, movie.ReleaseDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	//
	return id, nil
}

func (r *MoviePG) AddActor(movieID int, actorID int) error {
	query := fmt.Sprintf("INSERT INTO actors_movies (actor_id, movie_id) values ($1, $2)")
	err := r.db.QueryRow(query, actorID, movieID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *MoviePG) UpdateTitle(movie entity.Movies) error {
	query := fmt.Sprintf("UPDATE movies SET title=$1 WHERE id=$2")
	_, err := r.db.Exec(query, movie.Title, movie.ID)
	return err
}

func (r *MoviePG) UpdateDescription(movie entity.Movies) error {
	query := fmt.Sprintf("UPDATE movies SET description=$1 WHERE id=$2")
	_, err := r.db.Exec(query, movie.Description, movie.ID)
	return err
}

func (r *MoviePG) UpdateRating(movie entity.Movies) error {
	query := fmt.Sprintf("UPDATE movies SET rating=$1 WHERE id=$2")
	_, err := r.db.Exec(query, movie.Rating, movie.ID)
	return err
}

func (r *MoviePG) UpdateReleaseDate(movie entity.Movies) error {
	query := fmt.Sprintf("UPDATE movies SET release_date=$1 WHERE id=$2")
	_, err := r.db.Exec(query, movie.ReleaseDate.GoString(), movie.ID)
	return err
}

func (r *MoviePG) DeleteActors(movieID int) error {
	query := fmt.Sprintf("DELETE FROM actors_movies WHERE movie_id=$1")
	_, err := r.db.Exec(query, movieID)
	return err
}

func (r *MoviePG) Delete(movieID int) error {
	err := r.DeleteActors(movieID)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM movies WHERE id=$1")
	_, err = r.db.Exec(query, movieID)
	return err
}
