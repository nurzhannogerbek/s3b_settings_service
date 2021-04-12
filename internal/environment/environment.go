package environment

import "os"

var (
	PostgreSQLUserName              = os.Getenv("POSTGRESQL_USERNAME")
	PostgreSQLPassword              = os.Getenv("POSTGRESQL_PASSWORD")
	PostgreSQLHost                  = os.Getenv("POSTGRESQL_HOST")
	PostgreSQLPort                  = os.Getenv("POSTGRESQL_PORT")
	PostgreSQLDBName                = os.Getenv("POSTGRESQL_DB_NAME")
	TelegramBotURL                  = os.Getenv("TELEGRAM_BOT_URL")
	FacebookMessengerBotURL         = os.Getenv("FACEBOOK_MESSENGER_BOT_URL")
	FacebookMessengerBotVerifyToken = os.Getenv("FACEBOOK_MESSENGER_BOT_VERIFY_TOKEN")
	FacebookAppId                   = os.Getenv("FACEBOOK_APP_ID")
	FacebookAppSecret               = os.Getenv("FACEBOOK_APP_SECRET")
	InstagramPrivateBotURL          = os.Getenv("INSTAGRAM_PRIVATE_BOT_URL")
)