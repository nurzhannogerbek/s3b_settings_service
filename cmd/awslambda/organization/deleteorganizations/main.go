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

type OrganizationEvent struct {
	OrganizationsIDs *[]string `json:"organizationsIds"`
}

type OrganizationResponse struct {
	OrganizationsIDs *[]string `json:"organizationsIds"`
	Deleted          bool      `json:"deleted"`
}

func handleRequest(e common.Event) (interface{}, error) {
	var organizationEvent OrganizationEvent
	if err := json.Unmarshal(e.Arguments, &organizationEvent); err != nil {
		return nil, err
	}

	err := Services.Organization.DeleteOrganizations(organizationEvent.OrganizationsIDs)
	if err != nil {
		return nil, err
	}

	organizationResponse := OrganizationResponse{
		OrganizationsIDs: organizationEvent.OrganizationsIDs,
		Deleted:          true,
	}

	return organizationResponse, nil
}

func main() {
	lambda.Start(handleRequest)
}
