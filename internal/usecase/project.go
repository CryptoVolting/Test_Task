package usecase

import (
	"testProject/internal/repository"
	"testProject/pkg"
)

type ProjectUsecase struct {
	repo repository.Proj
}

func NewProjectUsecase(repo repository.Proj) *ProjectUsecase {
	return &ProjectUsecase{repo: repo}
}

func (s *ProjectUsecase) Create(project pkg.Project) (string, error) {
	return s.repo.Create(project)
}

func (s *ProjectUsecase) GetAll() ([]pkg.Project, error) {
	return s.repo.GetAll()
}

func (s *ProjectUsecase) GetById(id string) (pkg.Project, error) {
	return s.repo.GetById(id)
}

func (s *ProjectUsecase) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *ProjectUsecase) UpdateById(id string, projectUpdate pkg.UpdateProjectInput) error {
	if err := projectUpdate.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(id, projectUpdate)
}

func (s *ProjectUsecase) CreateAssign(input pkg.IdOperatorAndProject) (string, error) {
	return s.repo.CreateAssign(input)
}

func (s *ProjectUsecase) DeleteByIdAssign(id int) error {
	return s.repo.DeleteByIdAssign(id)
}
