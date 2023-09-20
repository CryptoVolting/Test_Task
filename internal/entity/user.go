package entity

import "errors"

type User struct {
	Id       int    `json:"-"        db:"id"`
	Name     string `json:"name"     binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Admin    *bool  `json:"admin"    binding:"required" db:"is_admin"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    *bool  `json:"admin"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == "" && i.Username == "" && i.Password == "" && i.Admin == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
