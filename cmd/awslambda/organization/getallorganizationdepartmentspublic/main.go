package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
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

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryString, ok := request.QueryStringParameters["rootOrganizationId"]
	if ok != true {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "key does not exist",
			IsBase64Encoded:   false,
		}, nil
	}

	if err := uuid.Validate(&queryString); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              err.Error(),
			IsBase64Encoded:   false,
		}, nil
	}

	organizations, err := Services.Organization.GetAllOrganizationDepartments(&queryString)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              err.Error(),
			IsBase64Encoded:   false,
		}, err
	}

	org, err := json.Marshal(&organizations)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              err.Error(),
			IsBase64Encoded:   false,
		}, err
	}

	return events.APIGatewayProxyResponse{Body: string(org), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
