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

func handleRequest(e common.Event) (interface{}, error) {
	var organizationSettingsEvent OrganizationSettingsEvent
	if err := json.Unmarshal(e.Arguments, &organizationSettingsEvent); err != nil {
		return nil, err
	}

	organizationSettings, err := Services.OrganizationSettings.GetOrganizationSettingsByID(&organizationSettingsEvent.OrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}

func main() {
	lambda.Start(handleRequest)
}