package usecase

import (
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
)

type OperatorUsage interface {
	Create(operator entity.Operator) (string, error)
	GetAll() ([]entity.Operator, error)
	GetById(id string) (entity.Operator, error)
	UpdateById(id string, operatorUpdate entity.UpdateOperatorInput) error
	DeleteById(id string) error
}

type ProjectUsage interface {
	Create(project entity.Project) (string, error)
	GetAll() ([]entity.Project, error)
	GetById(id string) (entity.Project, error)
	UpdateById(id string, projectUpdate entity.UpdateProjectInput) error
	DeleteById(id string) error
	CreateAssign(input entity.IdOperatorAndProject) (string, error)
	DeleteByIdAssign(id int) error
}

type Usecase struct {
	OperatorUsage
	ProjectUsage
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		OperatorUsage: NewOperatorUsecase(repository.OperatorUsage),
		ProjectUsage:  NewProjectUsecase(repository.ProjectUsage),
	}
}
