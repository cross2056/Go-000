package data

import (
	"database/sql"
	"fmt"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (repo *MessageRepo) GetMore(name string) string {
	// query db
	// query := "select more from table where name = msg"
	// more := repo.db.Exec(query)
	return fmt.Sprintf("this is result from DB for Name: %s", name)
}
