package repository

import (
	"github.com/jmoiron/sqlx"
	"testProject/internal/entity"
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

type Repository struct {
	OperatorUsage
	ProjectUsage
}

func NewSRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OperatorUsage: NewOperatorPostgres(db),
		ProjectUsage:  NewProjectPostgres(db),
	}
}
