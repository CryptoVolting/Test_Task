package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"testProject/internal/entity"
)

type PanelPostgres struct {
	db *sqlx.DB
}

func NewPanelPostgres(db *sqlx.DB) *PanelPostgres {
	return &PanelPostgres{db: db}
}

func (r *PanelPostgres) CreateUser(user entity.User) (int, error) {
	var (
		id     int
		roleID string
	)
	query := fmt.Sprintf("INSERT INTO%s (name, username, password_hash) values ($1, $2, $3) RETURNING id, admin_id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id, &roleID); err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO%s (id, is_admin) values ($1, $2)", rolesTable)
	r.db.QueryRow(query2, roleID, user.Admin)

	query3 := fmt.Sprintf("INSERT INTO%s (permission_name) values ('/panel/admin/%d')", adminPermissons, id)
	r.db.QueryRow(query3)

	return id, nil
}

func (r *PanelPostgres) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM%s where id='%s';", usersTable, id)
	_, err := r.db.Exec(query)

	query2 := fmt.Sprintf("DELETE FROM%s where permission_name='/panel/admin/%s';", adminPermissons, id)
	_, err = r.db.Exec(query2)

	return err
}

func (r *PanelPostgres) GetAll() ([]entity.User, error) {
	var lists []entity.User

	query := fmt.Sprintf("SELECT users.id, users.name, users.username, users.password_hash, roles.is_admin FROM%s,%s "+
		"WHERE users.admin_id = roles.id;", usersTable, rolesTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *PanelPostgres) UpdateById(userId string, userUpdate entity.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if userUpdate.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, userUpdate.Name)
		argId++
	}

	if userUpdate.Username != "" {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, userUpdate.Username)
		argId++
	}

	if userUpdate.Password != "" {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, userUpdate.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE%s SET %s WHERE id=$%d;", usersTable, setQuery, argId)
	args = append(args, userId)

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	var adminID string
	query2 := fmt.Sprintf("SELECT admin_id FROM%s WHERE id=$1", usersTable)
	if err := r.db.Get(&adminID, query2, userId); err != nil {
		return err
	}

	query3 := fmt.Sprintf("UPDATE%s SET is_admin='%t' WHERE id='%s';", rolesTable, *userUpdate.Admin, adminID)
	if _, err := r.db.Exec(query3); err != nil {
		return err
	}
	return nil
}
