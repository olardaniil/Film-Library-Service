package entity

type User struct {
	ID           int    `json:"id,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	Password     string `json:"password,omitempty"`
	HashPassword string `json:"hash_password,omitempty"`
	IsAdmin      bool   `json:"is_admin,omitempty"`
	Token        string `json:"token,omitempty"`
}
