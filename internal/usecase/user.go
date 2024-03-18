package usecase

import (
	"film-library-service/internal/repository"
	"log"
)

type UserUsecase struct {
	repo repository.UserRepo
}

func NewUserUsecase(repo repository.UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u UserUsecase) IsUserAdmin(userID int) bool {
	err := u.repo.IsUserAdmin(userID)
	if err != nil {
		if err.Error() != "no rows in result set" {
			log.Println(err)
		}
		return false
	}
	return true
}
