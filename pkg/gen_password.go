package pkg

import (
	"github.com/sethvargo/go-password/password"
	"testProject/internal/entity"
)

func NewPassword(operator *entity.Operator) error {
	res, err := password.Generate(15, 3, 3, false, false)
	if err != nil {
		return err
	}
	operator.Password = res

	return nil
}
