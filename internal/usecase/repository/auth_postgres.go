package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testProject/internal/entity"
)

const (
	usersTable = " users"
	rolesTable = " roles"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
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

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	var adminID string
	query := fmt.Sprintf("SELECT admin_id FROM%s WHERE username=$1 AND password_hash=$2", usersTable)
	if err := r.db.Get(&adminID, query, username, password); err != nil {
		return user, err
	}

	query2 := fmt.Sprintf("SELECT is_admin FROM%s WHERE id=$1", rolesTable)
	if err := r.db.Get(&user.Admin, query2, adminID); err != nil {
		return user, err
	}

	return user, nil
}
