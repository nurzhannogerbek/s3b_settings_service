package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Connect *sqlx.DB
}

func NewClient(connString *string) Connection {
	db, err := sqlx.Connect("pgx", *connString)
	if err != nil {
		panic(err)
	}
	return Connection{
		Connect: db,
	}
}

func (c *Connection) Ping() error {
	if err := c.Connect.Ping(); err != nil {
		return err
	}
	return nil
}