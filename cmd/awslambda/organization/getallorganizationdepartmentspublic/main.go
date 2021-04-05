package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"encoding/json"
	"fmt"
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

func createResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        statusCode,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              body,
		IsBase64Encoded:   false,
	}
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Headers)
	token, ok := request.Headers["authorization"]
	if ok != true {
		return createResponse(400, "Unauthorized"), nil
	}

	if token != "Bearer c10321f9-4574-47bc-a7b6-f101770dbd97" {
		return createResponse(400, "Invalid Token"), nil
	}

	queryString, ok := request.QueryStringParameters["rootOrganizationId"]
	if ok != true {
		return createResponse(400, fmt.Sprintf("query string %+v does not exist, only rootOrganizationId", request.QueryStringParameters)), nil
	}

	organizations, err := Services.Organization.GetAllOrganizationDepartments(&queryString)
	if err != nil {
		return createResponse(400, err.Error()), nil
	}

	org, _ := json.Marshal(&organizations)

	return createResponse(200, string(org)), nil
}

func main() {
	lambda.Start(handleRequest)
}
