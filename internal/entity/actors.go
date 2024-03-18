package entity

import "time"

type Actors struct {
	ID         int       `json:"id,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	MiddleName string    `json:"middle_name,omitempty"`
	GenderID   int       `json:"gender_id,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	DateBirth  time.Time `json:"date_birth"`
	Movie      []Movies  `json:"movie,omitempty"`
}

type ActorsInputBody struct {
	LastName   string    `json:"last_name,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	MiddleName string    `json:"middle_name,omitempty"`
	GenderID   int       `json:"gender_id,omitempty"`
	DateBirth  time.Time `json:"date_birth"`
}

type Gender struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var Genders []Gender = []Gender{
	{ID: 1, Name: "Мужчина"},
	{ID: 2, Name: "Женщина"},
}

func (a *Actors) ValidationLastName() bool {
	if len(a.LastName) < 2 {
		return false
	}
	return true
}

func (a *Actors) ValidationFirstName() bool {
	if len(a.FirstName) < 2 {
		return false
	}
	return true
}

func (a *Actors) ValidationGender() bool {
	for _, g := range Genders {
		if g.ID == a.GenderID {
			return true
		}
	}
	return false
}

func (a *Actors) SetGenderName() {
	for _, g := range Genders {
		if g.ID == a.GenderID {
			a.Gender = g.Name
		}
	}

}
