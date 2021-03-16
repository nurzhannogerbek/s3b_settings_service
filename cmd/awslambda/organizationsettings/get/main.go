package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	postgresqlrepo "bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository/postgrsql"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
)

var organizationService *service.Services

func init() {
	config, err := postgresql.NewConfig(
		environment.PostgreSQLUserName,
		environment.PostgreSQLPassword,
		environment.PostgreSQLHost,
		environment.PostgreSQLPort,
		environment.PostgreSQLDBName,
		"require",
	)
	if err != nil {
		panic(err)
	}

	connString, err := config.GetConnectionString()
	if err != nil {
		panic(err)
	}

	db := postgresql.NewClient(connString)
	organizationRepository := postgresqlrepo.NewRepositories(db)
	organizationService = service.NewServices(service.Dependencies{Repositories: organizationRepository})
}

type OrganizationSettingsEvent struct {
	OrganizationID string `json:"organizationId"`
}

func handleRequest(event common.Event) (interface{}, error) {
	var organizationSettingsEvent OrganizationSettingsEvent
	if err := json.Unmarshal(event.Arguments, &organizationSettingsEvent); err != nil {
		return nil, err
	}

	organizationSettings, err := organizationService.OrganizationSettings.Get(&organizationSettingsEvent.OrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}

func main() {
	lambda.Start(handleRequest)
}