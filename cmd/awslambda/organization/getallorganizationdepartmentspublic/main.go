package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"encoding/json"
	"errors"
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
	token, ok := request.Headers["Authorization"]
	if ok != true {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, errors.New("Unauthorized")
	}

	if token != "c10321f9-4574-47bc-a7b6-f101770dbd97" {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, errors.New("Invalid Token")
	}

	queryString, ok := request.QueryStringParameters["rootOrganizationId"]
	if ok != true {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, errors.New("key does not exist")
	}

	if err := uuid.Validate(&queryString); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, err
	}

	organizations, err := Services.Organization.GetAllOrganizationDepartments(&queryString)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, err
	}

	org, err := json.Marshal(&organizations)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: nil,
		MultiValueHeaders: nil,
		Body: string(org),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
