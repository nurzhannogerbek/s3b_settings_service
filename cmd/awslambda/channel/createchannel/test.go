package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	client := &http.Client{}

	data := url.Values{}
	data.Set("channel_id", "1e49e1ca-e57c-41b1-86ab-182df4753ee5")
	data.Set("action", "false")

	urlAddress := "https://instagram-private-service.3beep.io/dev/toggle_session_for_instagram_private_channel"

	request, err := http.NewRequest("POST", urlAddress, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
