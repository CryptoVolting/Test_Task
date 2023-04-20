package service

import (
	"testProject"
	"testProject/pkg/repository"
)

type OperatorService struct {
	repo repository.Oper
}

func NewOperatorService(repo repository.Oper) *OperatorService {
	return &OperatorService{repo: repo}
}

func (s *OperatorService) Create(operator testProject.Operator) (string, error) {
	return s.repo.Create(operator)
}

func (s *OperatorService) GetAll() ([]testProject.Operator, error) {
	return s.repo.GetAll()
}

func (s *OperatorService) GetById(id string) (testProject.Operator, error) {
	return s.repo.GetById(id)
}

func (s *OperatorService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *OperatorService) UpdateById(id string, operatorUpdate testProject.UpdateOperatorInput) error {
	if err := operatorUpdate.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(id, operatorUpdate)
}
