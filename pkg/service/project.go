package service

import (
	"testProject"
	"testProject/pkg/repository"
)

type ProjectService struct {
	repo repository.Proj
}

func NewProjectService(repo repository.Proj) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create(project testProject.Project) (string, error) {
	return s.repo.Create(project)
}

func (s *ProjectService) GetAll() ([]testProject.Project, error) {
	return s.repo.GetAll()
}

func (s *ProjectService) GetById(id string) (testProject.Project, error) {
	return s.repo.GetById(id)
}

func (s *ProjectService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *ProjectService) UpdateById(id string, projectUpdate testProject.UpdateProjectInput) error {
	if err := projectUpdate.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(id, projectUpdate)
}

func (s *ProjectService) CreateAssign(input testProject.IdOperatorAndProject) (string, error) {
	return s.repo.CreateAssign(input)
}

func (s *ProjectService) DeleteByIdAssign(id int) error {
	return s.repo.DeleteByIdAssign(id)
}
