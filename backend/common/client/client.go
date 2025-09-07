package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func DoRequest(client *http.Client, url string) ([]byte, error) {
	// Hitting API With Client GET
	resp, err := client.Get(url)

	// Checking Error
	if err != nil {
		return nil, err
	}
	

	// Close Client Conn
	defer resp.Body.Close()

	// Handling Data Response
	// Checking Error Response
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("response error with response code not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// If Status Code OK Then Decode The Response
	return body, nil
}
