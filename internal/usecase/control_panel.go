package usecase

import (
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
)

type PanelUsecase struct {
	panelUsage repository.PanelUsage
}

func NewPanelUsecase(panelUsage repository.PanelUsage) *PanelUsecase {
	return &PanelUsecase{panelUsage: panelUsage}
}

func (s *PanelUsecase) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.panelUsage.CreateUser(user)
}

func (s *PanelUsecase) UpdateById(id string, updateUser entity.UpdateUserInput) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}
	updateUser.Password = generatePasswordHash(updateUser.Password)
	return s.panelUsage.UpdateById(id, updateUser)
}

func (s *PanelUsecase) DeleteById(id string) error {
	return s.panelUsage.DeleteById(id)
}

func (s *PanelUsecase) GetAll() ([]entity.User, error) {
	return s.panelUsage.GetAll()
}
