package usecase

import (
	"film-library-service/internal/entity"
	"film-library-service/internal/repository"
	"fmt"
	"strconv"
	"time"
)

type ActorUsecase struct {
	repo repository.ActorRepo
}

func NewActorUsecase(repo repository.ActorRepo) *ActorUsecase {
	return &ActorUsecase{repo: repo}
}

func (u *ActorUsecase) GetAll() ([]entity.Actors, error) {
	actors, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, actor := range actors {
		moviesActors, err := u.repo.GetMoviesByActorID(actor.ID)
		if err != nil {
			return nil, err
		}

		for _, movies := range moviesActors {
			actor.Movie = append(actor.Movie, movies)
			actors[i] = actor
		}

	}

	return actors, nil
}

func (u *ActorUsecase) Create(actor entity.Actors) (int, error) {
	id, err := u.repo.Create(actor)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *ActorUsecase) Update(actorID int, lName string, fName string, mName string, genderID string, dateBirth string) error {

	if lName != "" {
		err := u.repo.UpdateLastName(entity.Actors{
			ID:       actorID,
			LastName: lName,
		})
		if err != nil {
			return err
		}
	}
	if fName != "" {
		err := u.repo.UpdateFirstName(entity.Actors{
			ID:        actorID,
			FirstName: fName,
		})
		if err != nil {
			return err
		}
	}
	if mName != "" {
		err := u.repo.UpdateMiddleName(entity.Actors{
			ID:         actorID,
			MiddleName: mName,
		})
		if err != nil {
			return err
		}
	}
	if genderID != "" {
		gID, err := strconv.Atoi(genderID)
		if err != nil {
			return fmt.Errorf("Ошибка при конвертирования строки в число: ", err)
		}
		err = u.repo.UpdateGenderID(entity.Actors{
			ID:       actorID,
			GenderID: gID,
		})
		if err != nil {
			return err
		}
	}
	if dateBirth != "" {
		dBirth, err := time.Parse("2006-01-02", dateBirth)
		if err != nil {
			return fmt.Errorf("Ошибка при конвертирования строки в дату: ", err)
		}
		err = u.repo.UpdateDateBirth(entity.Actors{
			ID:        actorID,
			DateBirth: dBirth,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *ActorUsecase) Delete(actorID int) error {
	err := u.repo.Delete(actorID)
	if err != nil {
		return err
	}

	return nil
}
