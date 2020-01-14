package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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


	// Open a direct message//////////////////////////////////////////////////////
	url := "https://slack.com/api/users.info"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User", selectedID)
	req.Header.Add("Authorization", "Bearer " + apiKey)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	/////////////////////////////////////////////////////////////////////////////////





	// Send a direct message/////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////



	fmt.Println(body)
}