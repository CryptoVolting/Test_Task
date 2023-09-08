package repository

import (
	"github.com/jmoiron/sqlx"
	"testProject/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

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
	Authorization
	OperatorUsage
	ProjectUsage
}

func NewSRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		OperatorUsage: NewOperatorPostgres(db),
		ProjectUsage:  NewProjectPostgres(db),
	}
}
