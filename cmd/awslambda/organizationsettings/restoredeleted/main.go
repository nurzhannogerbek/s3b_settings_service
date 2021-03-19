package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
)

var Services *service.Services

func init() {
	config := postgresql.NewConfig(
		environment.PostgreSQLUserName,
		environment.PostgreSQLPassword,
		environment.PostgreSQLHost,
		environment.PostgreSQLPort,
		environment.PostgreSQLDBName,
		"require",
	)

	postgresqlDB := postgresql.NewClient(config.GetConnectionString())
	repositories := repository.NewRepositories()
	repositories.SetPostgresqlRepositories(postgresqlDB)

	Services = service.NewServices(service.Dependencies{Repositories: repositories})
}

type OrganizationSettingsEvent struct {
	OrganizationID string `json:"organizationId"`
}

type OrganizationSettingsResponse struct {
	OrganizationSettingsEvent
	Restored bool `json:"restored"`
}

func handleRequest(e common.Event) (interface{}, error) {
	var organizationSettingsEvent OrganizationSettingsEvent
	if err := json.Unmarshal(e.Arguments, &organizationSettingsEvent); err != nil {
		return nil, err
	}

	if err := Services.OrganizationSettings.RestoreDeleted(&organizationSettingsEvent.OrganizationID); err != nil {
		return nil, err
	}

	organizationSettingsResponse := OrganizationSettingsResponse{
		OrganizationSettingsEvent: organizationSettingsEvent,
		Restored:                   true,
	}

	return organizationSettingsResponse, nil
}

func main() {
	lambda.Start(handleRequest)
}
