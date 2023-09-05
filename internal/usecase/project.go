package usecase

import (
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
)

type ProjectUsecase struct {
	projectUsage repository.ProjectUsage
}

func NewProjectUsecase(projectUsage repository.ProjectUsage) *ProjectUsecase {
	return &ProjectUsecase{projectUsage: projectUsage}
}

func (s *ProjectUsecase) Create(project entity.Project) (string, error) {
	return s.projectUsage.Create(project)
}

func (s *ProjectUsecase) GetAll() ([]entity.Project, error) {
	return s.projectUsage.GetAll()
}

func (s *ProjectUsecase) GetById(id string) (entity.Project, error) {
	return s.projectUsage.GetById(id)
}

func (s *ProjectUsecase) DeleteById(id string) error {
	return s.projectUsage.DeleteById(id)
}

func (s *ProjectUsecase) UpdateById(id string, projectUpdate entity.UpdateProjectInput) error {
	if err := projectUpdate.Validate(); err != nil {
		return err
	}
	return s.projectUsage.UpdateById(id, projectUpdate)
}

func (s *ProjectUsecase) CreateAssign(input entity.IdOperatorAndProject) (string, error) {
	return s.projectUsage.CreateAssign(input)
}

func (s *ProjectUsecase) DeleteByIdAssign(id int) error {
	return s.projectUsage.DeleteByIdAssign(id)
}
