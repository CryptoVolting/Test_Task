package usecase

import (
	"testProject/internal/repository"
	"testProject/pkg"
)

type OperatorUsecase struct {
	repo repository.Oper
}

func NewOperatorUsecase(repo repository.Oper) *OperatorUsecase {
	return &OperatorUsecase{repo: repo}
}

func (s *OperatorUsecase) Create(operator pkg.Operator) (string, error) {
	return s.repo.Create(operator)
}

func (s *OperatorUsecase) GetAll() ([]pkg.Operator, error) {
	return s.repo.GetAll()
}

func (s *OperatorUsecase) GetById(id string) (pkg.Operator, error) {
	return s.repo.GetById(id)
}

func (s *OperatorUsecase) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *OperatorUsecase) UpdateById(id string, operatorUpdate pkg.UpdateOperatorInput) error {
	if err := operatorUpdate.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(id, operatorUpdate)
}
