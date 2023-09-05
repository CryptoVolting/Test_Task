package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"testProject/internal/entity"
)

const (
	operatorsTable    = " operators"
	projectTable      = " projects"
	projOperListTable = " project_operators_list"
)

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres(db *sqlx.DB) *ProjectPostgres {
	return &ProjectPostgres{db: db}
}

func (r *ProjectPostgres) Create(project entity.Project) (string, error) {
	var id string
	createProjectQuery := fmt.Sprintf("INSERT INTO%s (name, typeProject) VALUES ($1, $2) RETURNING id", projectTable)
	row := r.db.QueryRow(createProjectQuery, project.Name, project.TypeProject)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProjectPostgres) GetAll() ([]entity.Project, error) {
	var lists []entity.Project

	query := fmt.Sprintf("SELECT * FROM%s;", projectTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *ProjectPostgres) GetById(id string) (entity.Project, error) {
	var project entity.Project

	query := fmt.Sprintf("SELECT * FROM%s where id='%s';", projectTable, id)
	err := r.db.Get(&project, query)

	return project, err
}

func (r *ProjectPostgres) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM%s where id='%s';", projectTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *ProjectPostgres) UpdateById(projectId string, projectUpdate entity.UpdateProjectInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if projectUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *projectUpdate.Name)
		argId++
	}

	if projectUpdate.TypeProject != nil {
		setValues = append(setValues, fmt.Sprintf("typeProject=$%d", argId))
		args = append(args, *projectUpdate.TypeProject)
		argId++
	}

	// name=$1
	// typeProject=$1
	// name=$1, typeProject=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE%s SET %s WHERE id=$%d;", projectTable, setQuery, argId)
	args = append(args, projectId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ProjectPostgres) CreateAssign(input entity.IdOperatorAndProject) (string, error) {
	var id string
	createProjectQuery := fmt.Sprintf("INSERT INTO%s (operators_id, project_id) VALUES ($1, $2) RETURNING id", projOperListTable)
	row := r.db.QueryRow(createProjectQuery, input.IdOperator, input.IdProject)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProjectPostgres) DeleteByIdAssign(id int) error {
	query := fmt.Sprintf("DELETE FROM%s where id=%d;", projOperListTable, id)
	_, err := r.db.Exec(query)

	return err
}
