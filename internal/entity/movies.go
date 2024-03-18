package entity

import (
	"fmt"
	"time"
)

type Movies struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Rating      float32   `json:"rating,omitempty"`
	ActorIDs    []int     `json:"actor_ids,omitempty"`
}

type MoviesInputBody struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Rating      float32   `json:"rating,omitempty"`
	ActorIDs    []int     `json:"actor_ids,omitempty"`
}

func (m *Movies) Validation() error {
	if len(m.Title) < 1 || len(m.Title) > 150 {
		return fmt.Errorf("Название фильма должно содержать не менее 1 и не более 150 символов")
	}
	if len(m.Description) > 1000 {
		return fmt.Errorf("Описание фильма должно содержать не более 1000 символов")
	}
	if m.Rating < 0 || m.Rating > 10 {
		return fmt.Errorf("Рейтинг фильма должен быть от 0 до 10")
	}

	return nil
}

func (m *Movies) ValidationForUpdate() error {
	if m.ID == 0 {
		return fmt.Errorf("Не указан ID фильма")
	}
	if len(m.Title) > 150 {
		return fmt.Errorf("Название фильма должно содержать не менее 1 и не более 150 символов")
	}
	if len(m.Description) > 1000 {
		return fmt.Errorf("Описание фильма должно содержать не более 1000 символов")
	}
	if m.Rating < 0 || m.Rating > 10 {
		return fmt.Errorf("Рейтинг фильма должен быть от 0 до 10")
	}

	return nil
}

func (m *Movies) ValidationForDelete() error {
	if m.ID == 0 {
		return fmt.Errorf("Не указан ID фильма")
	}

	return nil
}
