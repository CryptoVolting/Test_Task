package usecase

import (
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
)

type OperatorUsecase struct {
	operatorUsage repository.OperatorUsage
}

func NewOperatorUsecase(operatorUsage repository.OperatorUsage) *OperatorUsecase {
	return &OperatorUsecase{operatorUsage: operatorUsage}
}

func (s *OperatorUsecase) Create(operator entity.Operator) (string, error) {
	return s.operatorUsage.Create(operator)
}

func (s *OperatorUsecase) GetAll() ([]entity.Operator, error) {
	return s.operatorUsage.GetAll()
}

func (s *OperatorUsecase) GetById(id string) (entity.Operator, error) {
	return s.operatorUsage.GetById(id)
}

func (s *OperatorUsecase) DeleteById(id string) error {
	return s.operatorUsage.DeleteById(id)
}

func (s *OperatorUsecase) UpdateById(id string, operatorUpdate entity.UpdateOperatorInput) error {
	if err := operatorUpdate.Validate(); err != nil {
		return err
	}
	return s.operatorUsage.UpdateById(id, operatorUpdate)
}
