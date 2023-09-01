package usecase

import (
	"testProject/internal/repository"
	"testProject/pkg"
)

type Oper interface {
	Create(operator pkg.Operator) (string, error)
	GetAll() ([]pkg.Operator, error)
	GetById(id string) (pkg.Operator, error)
	UpdateById(id string, operatorUpdate pkg.UpdateOperatorInput) error
	DeleteById(id string) error
}

type Proj interface {
	Create(project pkg.Project) (string, error)
	GetAll() ([]pkg.Project, error)
	GetById(id string) (pkg.Project, error)
	UpdateById(id string, projectUpdate pkg.UpdateProjectInput) error
	DeleteById(id string) error
	CreateAssign(input pkg.IdOperatorAndProject) (string, error)
	DeleteByIdAssign(id int) error
}

type Usecase struct {
	Oper
	Proj
}

func NewUsecase(repos *repository.Repository) *Usecase {
	return &Usecase{
		Oper: NewOperatorUsecase(repos.Oper),
		Proj: NewProjectUsecase(repos.Proj),
	}
}
