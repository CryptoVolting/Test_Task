package service

import (
	"testProject"
	"testProject/pkg/repository"
)

type Oper interface {
	Create(operator testProject.Operator) (string, error)
	GetAll() ([]testProject.Operator, error)
	GetById(id string) (testProject.Operator, error)
	UpdateById(id string, operatorUpdate testProject.UpdateOperatorInput) error
	DeleteById(id string) error
}

type Proj interface {
	Create(project testProject.Project) (string, error)
	GetAll() ([]testProject.Project, error)
	GetById(id string) (testProject.Project, error)
	UpdateById(id string, projectUpdate testProject.UpdateProjectInput) error
	DeleteById(id string) error
	CreateAssign(input testProject.IdOperatorAndProject) (string, error)
	DeleteByIdAssign(id int) error
}

type Service struct {
	Oper
	Proj
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Oper: NewOperatorService(repos.Oper),
		Proj: NewProjectService(repos.Proj),
	}
}
