package model

import (
	"database/sql"

	"github.com/pkg/errors"
)

// User struct
type User struct {
	ID   int
	Name string
}

// QueryUsers returns users data and error
func QueryUsers() ([]User, error) {
	data, err := queryDBWithError("select * from users")
	if err != nil {
		return nil, errors.Wrap(err, "no user data")
	}
	return data.([]User), nil
}

// queryDBWithError returns sql.ErrNoRows
func queryDBWithError(_ string) (interface{}, error) {
	// some sql query
	return nil, sql.ErrNoRows
}
