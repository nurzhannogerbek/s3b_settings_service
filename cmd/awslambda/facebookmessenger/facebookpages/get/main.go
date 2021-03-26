package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

const facebookGraphApiUrl string = "https://graph.facebook.com/v9.0"

type FacebookMessengerSettingsEvent struct {
	UserId                    string `json:"userId"`
	ShortLivedUserAccessToken string `json:"shortLivedUserAccessToken"`
}

func handleRequest(e common.Event) (*common.FacebookPages, error) {
	var facebookMessengerSettingsEvent FacebookMessengerSettingsEvent
	if err := json.Unmarshal(e.Arguments, &facebookMessengerSettingsEvent); err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("grant_type", "fb_exchange_token")
	parameters.Add("client_id", environment.FacebookAppId)
	parameters.Add("client_secret", environment.FacebookAppSecret)
	parameters.Add("fb_exchange_token", facebookMessengerSettingsEvent.ShortLivedUserAccessToken)

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/oauth/access_token", facebookGraphApiUrl), strings.NewReader(parameters.Encode()))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
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
