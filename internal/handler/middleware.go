package handler

import (
	"github.com/sethvargo/go-password/password"
	"testProject/pkg"
)

func newPassword(operator *pkg.Operator) error {
	res, err := password.Generate(15, 3, 3, false, false)
	if err != nil {
		return err
	}
	operator.Password = res

	return nil
}
