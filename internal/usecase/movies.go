package usecase

import (
	"film-library-service/internal/entity"
	"film-library-service/internal/repository"
	"fmt"
	"log"
	"strconv"
	"time"
)

type MovieUsecase struct {
	repo repository.MovieRepo
}

func NewMovieUsecase(repo repository.MovieRepo) *MovieUsecase {
	return &MovieUsecase{repo: repo}
}

func (u *MovieUsecase) GetAll(sort string) ([]entity.Movies, error) {
	movies, err := u.repo.GetAll(sort)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (u *MovieUsecase) Search(userQuery string) ([]entity.Movies, error) {
	movies, err := u.repo.Search(userQuery)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (u *MovieUsecase) Create(movie entity.Movies) (int, error) {
	movieID, err := u.repo.Create(movie)
	if err != nil {
		return 0, err
	}

	if len(movie.ActorIDs) != 0 {
		for _, actorID := range movie.ActorIDs {
			err := u.repo.AddActor(movieID, actorID)
			if err != nil {
				log.Println("Ошибка при добавлении актёра: ", err.Error())
			}
		}
	}

	return movieID, nil
}

func (u *MovieUsecase) Update(movieID int, title string, description string, releaseDate string, rating string) error {
	//
	if title != "" {
		err := u.repo.UpdateTitle(entity.Movies{
			ID:    movieID,
			Title: title,
		})
		if err != nil {
			return err
		}
	}
	//
	if description != "" {
		err := u.repo.UpdateDescription(entity.Movies{
			ID:          movieID,
			Description: description,
		})
		if err != nil {
			return err
		}
	}
	//
	if releaseDate != "" {
		rDate, err := time.Parse("2006-01-02", releaseDate)
		if err != nil {
			return fmt.Errorf("Ошибка при конвертирования строки в дату: ", err)
		}
		err = u.repo.UpdateReleaseDate(entity.Movies{
			ID:          movieID,
			ReleaseDate: rDate,
		})
		if err != nil {
			return err
		}
	}
	//
	if rating != "" {
		r, err := strconv.Atoi(rating)
		if err != nil {
			return fmt.Errorf("Ошибка при конвертирования строки в число: ", err)
		}
		err = u.repo.UpdateRating(entity.Movies{
			ID:     movieID,
			Rating: float32(r),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *MovieUsecase) Delete(movieID int) error {
	err := u.repo.Delete(movieID)
	if err != nil {
		return err
	}

	return nil
}
