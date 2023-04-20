package testProject

import "errors"

type Operator struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Surname   string `json:"surname" db:"surname" binding:"required"`
	Town      string `json:"town" db:"town" binding:"required"`
	Telephone string `json:"telephone" db:"telephone" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"password"`
}

type UpdateOperatorInput struct {
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	Town      *string `json:"town"`
	Telephone *string `json:"telephone"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}

func (i UpdateOperatorInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Town == nil && i.Telephone == nil && i.Email == nil && i.Password == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
