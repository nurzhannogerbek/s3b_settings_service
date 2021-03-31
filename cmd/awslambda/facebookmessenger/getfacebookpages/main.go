package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
)

const facebookGraphApiUrl string = "https://graph.facebook.com/v9.0"

type Event struct {
	UserId                    string `json:"userId"`
	ShortLivedUserAccessToken string `json:"shortLivedUserAccessToken"`
}

type User struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func handleRequest(e common.Event) (*common.FacebookPages, error) {
	var event Event
	if err := json.Unmarshal(e.Arguments, &event); err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/oauth/access_token", facebookGraphApiUrl), nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	query.Add("grant_type", "fb_exchange_token")
	query.Add("client_id", environment.FacebookAppId)
	query.Add("client_secret", environment.FacebookAppSecret)
	query.Add("fb_exchange_token", event.ShortLivedUserAccessToken)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var user User
	var longLivedUserAccessToken string
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	longLivedUserAccessToken = user.AccessToken

	request, err = http.NewRequest("GET", fmt.Sprintf("%s/%s/accounts", facebookGraphApiUrl, event.UserId), nil)
	if err != nil {
		return nil, err
	}

	query = request.URL.Query()
	query.Add("access_token", longLivedUserAccessToken)
	request.URL.RawQuery = query.Encode()

	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var facebookPages *common.FacebookPages
	if err := json.Unmarshal(body, &facebookPages); err != nil {
		return nil, err
	}

	return facebookPages, nil
}

func main() {
	lambda.Start(handleRequest)
}
