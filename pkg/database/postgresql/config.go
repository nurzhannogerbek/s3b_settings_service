package postgresql

import (
	"errors"
	"fmt"
	"strconv"
)

type config struct {
	PostgreSQLUserName *string
	PostgreSQLPassword *string
	PostgreSQLHost     *string
	PostgreSQLPort     *int64
	PostgreSQLDBName   *string
	SSLMode            *string
}

// NewConfig - Creates new config of database, returns pointer to config and error
func NewConfig(userName, userPassword, host, port, dbName, sslMode string) (*config, error) {
	port16, err := strconv.ParseInt(port, 10, 16)
	if err != nil {
		return nil, err
	}

	newConfig := config{
		PostgreSQLUserName: &userName,
		PostgreSQLPassword: &userPassword,
		PostgreSQLHost:     &host,
		PostgreSQLPort:     &port16,
		PostgreSQLDBName:   &dbName,
		SSLMode:            &sslMode,
	}
	return &newConfig, nil
}

// GetConnectionString - Creates connection string from config struct and returns pointer to string
func (c *config) GetConnectionString() (*string, error) {
	if c.PostgreSQLUserName == nil {
		return nil, errors.New("Configure database (NewConfig) before calling GetConnectionString")
	}
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		*c.PostgreSQLUserName,
		*c.PostgreSQLPassword,
		*c.PostgreSQLHost,
		*c.PostgreSQLPort,
		*c.PostgreSQLDBName,
		*c.SSLMode,
	)
	return &connectionString, nil
}
