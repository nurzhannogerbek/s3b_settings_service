package database

type Database interface {
	NewClient(connString *string) error
}

type Connector interface {
	Ping() error
}

type Configer interface {
	GetConnectionString() (*string, error)
}

