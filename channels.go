/******************************************************************************************************
Get Channel names and IDs
*******************************************************************************************************/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type channel struct {
	ID   string
	Name string
}

type channels struct {
	Channels []channel
}

func main() {
	slackToken := "<INSERT TOKEN HERE>"
	slackEndpoint := "https://slack.com/api/"

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/conversations.list?types=public_channel,private_channel", slackEndpoint), nil)
	if err != nil {
		fmt.Printf("Error creating request %+v", err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", slackToken))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error getting Slack channels %+v", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body %+v", err)
		return
	}

	res := channels{}
	json.Unmarshal([]byte(body), &res)
	fmt.Println(res.Channels)
}
