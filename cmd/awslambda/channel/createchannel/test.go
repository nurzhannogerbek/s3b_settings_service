package main

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://graph.facebook.com/oauth/access_token", nil)
	if err != nil {
		fmt.Println(err)
	}

	query := request.URL.Query()
	query.Add("client_id", environment.FacebookAppId)
	query.Add("client_secret", environment.FacebookAppSecret)
	query.Add("grant_type", "client_credentials")
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	type FacebookAppToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	var facebookAppToken common.FacebookMessengerUser
	if err = json.Unmarshal(body, &facebookAppToken); err != nil {
		fmt.Println(err)
	}

	request, err = http.NewRequest("GET", fmt.Sprintf("https://graph.facebook.com/v9.0/%s/subscriptions", environment.FacebookAppId), nil)
	if err != nil {
		return err
	}

	query = request.URL.Query()
	query.Add("access_token", facebookAppToken.AccessToken)
	query.Add("object", "page")
	query.Add("callback_url", environment.FacebookMessengerBotURL)
	query.Add("fields", "messages,messaging_postbacks,messaging_optins,message_deliveries,message_reads,messaging_payments,messaging_pre_checkouts,messaging_checkout_updates,messaging_account_linking,messaging_referrals,message_echoes,messaging_game_plays,standby,messaging_handovers,messaging_policy_enforcement,message_reactions,inbox_labels")
	query.Add("include_values", "true")
	query.Add("verify_token", environment.FacebookMessengerBotVerifyToken)
	request.URL.RawQuery = query.Encode()

	response, err = client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("couldn't install webhook via Facebook Graph API, response status code: %q", response.StatusCode)
	}

	return nil
}
