package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/retgits/bitly/client"
)

type Request struct {
	endpoint string
	method   string
	payload  []byte
}
type Response struct {
	Link string `json:"link"`
}

var err error

// Main function implementing the bitly API
func generateLink(url string) (string, error) {

	token := os.Getenv("BITLY_API_KEY")
	if len(token) < 41 {
		err = fmt.Errorf("%s", "MISSING API KEY or wrong key")
		return "", err
	}
	bitly := client.NewClient().WithAccessToken(token) // Generate a new Bitly client

	if bitly == nil {
		err = fmt.Errorf("%s", "[*] Failed to create client")
		return "", err
	}

	buff := new(bytes.Buffer)
	data := map[string]string{
		"long_url": url, // just added long_url, as the other params are optional
	}

	fmt.Fprintf(buff, "{")
	for key, value := range data {
		fmt.Fprintf(buff, "\"%s\": \"%s\"", key, value)
	}
	fmt.Fprintf(buff, "}") //stupid way of constructing the call post body :O

	r := Request{ // creating payload
		endpoint: "/shorten",
		method:   "POST",
		payload:  buff.Bytes(),
	}

	resp, err := bitly.Call(r.endpoint, r.method, r.payload) // sending API call
	if err != nil {
		return "", err
	}

	var response Response

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return "", err
	}
	shortLink := fmt.Sprintf("%s", response.Link)
	return shortLink, nil
}

// helper function for exporting generateLink to C ( and eventually Python)
func exportLink(url string) {
	URL, err := checkURL(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		link, err := generateLink(URL)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(link)
		}
	}

}

// basic sanitizer for URLs, since this will mainly be used for shortening ngrok links also added a check to see if they are responding or not
func checkURL(url string) (string, error) {

	URL := path.Clean(url)
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s", "URL didn`t respond with 200 OK")
		return "", err
	}
	return URL, nil

}
