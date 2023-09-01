package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"testProject/pkg"
)

type OperatorPostgres struct {
	db *sqlx.DB
}

func NewOperatorPostgres(db *sqlx.DB) *OperatorPostgres {
	return &OperatorPostgres{db: db}
}

func (r *OperatorPostgres) Create(operator pkg.Operator) (string, error) {
	var id string
	createOperatorQuery := fmt.Sprintf("INSERT INTO%s (name, surname, town, telephone, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", operatorsTable)
	row := r.db.QueryRow(createOperatorQuery, operator.Name, operator.Surname, operator.Town, operator.Telephone, operator.Email, operator.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *OperatorPostgres) GetAll() ([]pkg.Operator, error) {
	var lists []pkg.Operator

	query := fmt.Sprintf("SELECT * FROM%s;", operatorsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *OperatorPostgres) GetById(id string) (pkg.Operator, error) {
	var operator pkg.Operator

	query := fmt.Sprintf("SELECT * FROM%s where id='%s';", operatorsTable, id)
	err := r.db.Get(&operator, query)

	return operator, err
}

func (r *OperatorPostgres) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM%s where id='%s';", operatorsTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *OperatorPostgres) UpdateById(operatorId string, operatorUpdate pkg.UpdateOperatorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if operatorUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *operatorUpdate.Name)
		argId++
	}

	if operatorUpdate.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *operatorUpdate.Surname)
		argId++
	}

	if operatorUpdate.Town != nil {
		setValues = append(setValues, fmt.Sprintf("town=$%d", argId))
		args = append(args, *operatorUpdate.Town)
		argId++
	}

	if operatorUpdate.Telephone != nil {
		setValues = append(setValues, fmt.Sprintf("telephone=$%d", argId))
		args = append(args, *operatorUpdate.Telephone)
		argId++
	}

	if operatorUpdate.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *operatorUpdate.Email)
		argId++
	}

	if operatorUpdate.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *operatorUpdate.Password)
		argId++
	}

	// name=$1
	// surname=$1
	// name=$1, surname=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE%s SET %s WHERE id=$%d;", operatorsTable, setQuery, argId)
	args = append(args, operatorId)

	_, err := r.db.Exec(query, args...)
	return err
}
