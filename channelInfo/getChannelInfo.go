package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

type slackStruct struct {
	Ok       bool `json:"ok"`
	Channels []struct {
		ID                      string        `json:"id"`
		Name                    string        `json:"name"`
		IsChannel               bool          `json:"is_channel"`
		IsGroup                 bool          `json:"is_group"`
		IsIm                    bool          `json:"is_im"`
		Created                 int           `json:"created"`
		IsArchived              bool          `json:"is_archived"`
		IsGeneral               bool          `json:"is_general"`
		Unlinked                int           `json:"unlinked"`
		NameNormalized          string        `json:"name_normalized"`
		IsShared                bool          `json:"is_shared"`
		ParentConversation      interface{}   `json:"parent_conversation"`
		Creator                 string        `json:"creator"`
		IsExtShared             bool          `json:"is_ext_shared"`
		IsOrgShared             bool          `json:"is_org_shared"`
		SharedTeamIds           []string      `json:"shared_team_ids"`
		PendingShared           []interface{} `json:"pending_shared"`
		PendingConnectedTeamIds []interface{} `json:"pending_connected_team_ids"`
		IsPendingExtShared      bool          `json:"is_pending_ext_shared"`
		IsMember                bool          `json:"is_member"`
		IsPrivate               bool          `json:"is_private"`
		IsMpim                  bool          `json:"is_mpim"`
		Topic                   struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []string `json:"previous_names"`
		NumMembers    int      `json:"num_members"`
	} `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type ChannelStruct struct {
	Name string
	ID string
}


func main() {

	apiKey := os.Getenv("api")
	if apiKey == "" {
		fmt.Println("API Key missing")
		os.Exit(1)
	}
	fmt.Println("The current API key used: ", apiKey)

	url := "https://slack.com/api/conversations.list"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	slackData := slackStruct{}

	err = json.Unmarshal(body, &slackData)
	if err != nil {
		fmt.Println(err)
	}


	var nameList []ChannelStruct
	for i := 0 ; i < len(slackData.Channels) ; i++ {
		Name := slackData.Channels[i].Name

		channelInfo := ChannelStruct{
			Name: Name,
			ID:   slackData.Channels[i].ID,
		}

		nameList = append(nameList, channelInfo)
	}

	fmt.Println(nameList)
}