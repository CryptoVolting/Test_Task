package usecase

import (
	"github.com/redis/go-redis/v9"
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
)

type Usecase struct {
	PanelUsage
	Authorization
	OperatorUsage
	ProjectUsage
	RedisUsage
}

func NewUsecase(repository *repository.Repository, cacheRepository *repository.RedisRepository, client *redis.Client) *Usecase {
	return &Usecase{
		PanelUsage:    NewPanelUsecase(repository.PanelUsage),
		Authorization: NewAuthUsecase(repository.Authorization),
		OperatorUsage: NewOperatorUsecase(repository.OperatorUsage),
		ProjectUsage:  NewProjectUsecase(repository.ProjectUsage),
		RedisUsage:    NewRedisUsecase(cacheRepository.RedisUsage, client),
	}
}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (*bool, error)
}

type PanelUsage interface {
	CreateUser(user entity.User) (int, error)
	GetAll() ([]entity.User, error)
	DeleteById(id string) error
	UpdateById(id string, userUpdate entity.UpdateUserInput) error
}

type RedisUsage interface {
	GetPermissionsByRole(isAdmin string) ([]entity.Premissoins, error)
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
