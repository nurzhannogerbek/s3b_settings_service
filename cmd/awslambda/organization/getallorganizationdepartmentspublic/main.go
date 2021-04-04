package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"

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
	RootOrganizationID *string `json:"rootOrganizationId"`
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var organizationEvent OrganizationEvent
	if err := json.Unmarshal([]byte(request.Body), &organizationEvent); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body: err.Error(),
		}, err
	}

	organizations, err := Services.Organization.GetAllOrganizationDepartments(organizationEvent.RootOrganizationID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body: err.Error(),
		}, err
	}

	org, err := json.Marshal(&organizations)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body: err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{Body: string(org), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
