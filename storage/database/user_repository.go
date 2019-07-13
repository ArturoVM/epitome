package database

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/arturovm/epitome"
	"github.com/arturovm/epitome/storage"
)

// UserRepository implements storage.UserRepository.
type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sql.DB) storage.UserRepository {
	dbx := sqlx.NewDb(db, "sqlite")
	return &UserRepository{
		db: dbx,
	}
}

// Add implements UserRepository.Add.
func (r *UserRepository) Add(user epitome.User) error {
	_, err := squirrel.Insert("users").
		Columns("username", "password", "salt").
		Values(user.Username,
			user.Credentials().Password,
			user.Credentials().Salt).
		RunWith(r.db).
		Exec()
	return err
}

// ByUsername implements UserRepository.ByUsername.
func (r *UserRepository) ByUsername(username string) (*epitome.User, error) {
	query, args, err := squirrel.Select("password", "salt").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "error building select query")
	}

	var credentials epitome.Credentials
	err = r.db.Get(&credentials, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error querying database")
	}

	user := epitome.NewUser(username, &credentials)
	return &user, nil
}
