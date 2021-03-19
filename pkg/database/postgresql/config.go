package postgresql

import (
	"fmt"
	"strconv"
)

// config
// Contains configuration data for postgresql.s
type config struct {
	PostgreSQLUserName *string
	PostgreSQLPassword *string
	PostgreSQLHost     *string
	PostgreSQLPort     *int64
	PostgreSQLDBName   *string
	SSLMode            *string
}

// NewConfig
// Creates new config of database, returns pointer to config or panics on error.
func NewConfig(userName, userPassword, host, port, dbName, sslMode string) *config {
	port16, err := strconv.ParseInt(port, 10, 16)
	if err != nil {
		panic(err)
	}

	newConfig := config{
		PostgreSQLUserName: &userName,
		PostgreSQLPassword: &userPassword,
		PostgreSQLHost:     &host,
		PostgreSQLPort:     &port16,
		PostgreSQLDBName:   &dbName,
		SSLMode:            &sslMode,
	}
	return &newConfig
}

// GetConnectionString
// Creates connection string from config struct and returns pointer to string.
func (c *config) GetConnectionString() *string {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		*c.PostgreSQLUserName,
		*c.PostgreSQLPassword,
		*c.PostgreSQLHost,
		*c.PostgreSQLPort,
		*c.PostgreSQLDBName,
		*c.SSLMode)
	return &connectionString
}
