package environment

import "os"

var (
	PostgreSQLUserName = os.Getenv("POSTGRESQL_USERNAME")
	PostgreSQLPassword = os.Getenv("POSTGRESQL_PASSWORD")
	PostgreSQLHost     = os.Getenv("POSTGRESQL_HOST")
	PostgreSQLPort     = os.Getenv("POSTGRESQL_PORT")
	PostgreSQLDBName   = os.Getenv("POSTGRESQL_DB_NAME")
	FacebookAppId      = os.Getenv("FACEBOOK_APP_ID")
	FacebookAppSecret  = os.Getenv("FACEBOOK_APP_SECRET")
)