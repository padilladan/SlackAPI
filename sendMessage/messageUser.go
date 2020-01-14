package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"bytes"
	"mime/multipart"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)

type conversationStruct struct {
	Ok          bool `json:"ok"`
	NoOp        bool `json:"no_op"`
	AlreadyOpen bool `json:"already_open"`
	Channel     struct {
		ID                 string      `json:"id"`
		Created            int         `json:"created"`
		IsArchived         bool        `json:"is_archived"`
		IsIm               bool        `json:"is_im"`
		IsOrgShared        bool        `json:"is_org_shared"`
		User               string      `json:"user"`
		LastRead           string      `json:"last_read"`
		Latest             interface{} `json:"latest"`
		UnreadCount        int         `json:"unread_count"`
		UnreadCountDisplay int         `json:"unread_count_display"`
		IsOpen             bool        `json:"is_open"`
		Priority           int         `json:"priority"`
	} `json:"channel"`
}

func main() {
	// Get API Key//////////////////////////////////////////////////////
	apiKey := os.Getenv("api")
	if apiKey == "" {
		fmt.Println("API Key missing")
		os.Exit(1)
	}
	////////////////////////////////////////////////////////////////////





	// Search for user//////////////////////////////////////////////////////
	fmt.Println("Use CB Tools to retrieve user ID")
	fmt.Println("Enter ID: ")
	ID := bufio.NewReader(os.Stdin)
	const inputdelimiter = '\n'
	selectedID, err := ID.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	selectedID = strings.Replace(selectedID, "\n", "", -1)
	////////////////////////////////////////////////////////////////////////





	//Open conversation with user//////////////////////////////////////////////////////////////////////////////////
	url := "https://slack.com/api/conversations.open"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("users", selectedID)
	_ = writer.WriteField("return_im", "true")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + apiKey)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	slackData := conversationStruct{}
	err = json.Unmarshal(body, &slackData)
	if err != nil {
		fmt.Println(err)
	}

	channelID := slackData.Channel.ID
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////




	//Prompt to send to user//////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Enter Text: ")
	text := bufio.NewReader(os.Stdin)
	const inputdelimiter2 = '\n'
	selectedtext, err := text.ReadString(inputdelimiter2)
	if err != nil {
		fmt.Println(err)
		return
	}
	selectedtext = strings.Replace(selectedtext, "\n", "", -1)
	if selectedtext == "" {
		fmt.Println("Cannot send an empty message")
		os.Exit(1)
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////////////






	//Send message to user//////////////////////////////////////////////////////////////////////////////////
	url2 := "https://slack.com/api/chat.postMessage"
	method2 := "POST"

	payload2 := &bytes.Buffer{}
	writer2 := multipart.NewWriter(payload2)
	_ = writer2.WriteField("channel", channelID)
	_ = writer2.WriteField("text", selectedtext)
	err2 := writer2.Close()
	if err2 != nil {
		fmt.Println(err2)
	}


	client2 := &http.Client {
	}
	req2, err := http.NewRequest(method2, url2, payload2)

	if err != nil {
		fmt.Println(err)
	}
	req2.Header.Add("Content-Type", "application/json")
	req2.Header.Add("Authorization", "Bearer " + apiKey)

	req2.Header.Set("Content-Type", writer2.FormDataContentType())
	res2, err := client2.Do(req2)
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res2.Body)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////




	fmt.Println("Message sent!")
}