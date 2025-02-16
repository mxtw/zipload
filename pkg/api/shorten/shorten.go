package shorten

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/mxtw/zipload/pkg/api"
)

type postBody struct {
	Url    string `json:"url"`
	Vanity string `json:"vanity"`
}

type shortenResponse struct {
	Url string `json:"url"`
}

func Shorten(client *api.Client, targetUrl string, vanity string, options api.Options) (string, error) {
	endpoint, err := url.JoinPath(client.Host, "/api/shorten")
	if err != nil {
		log.Println(err)
		return "", err
	}

	body := postBody{
		targetUrl,
		vanity,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	log.Println(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header = options.ToHeaders()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", client.Token)

	hc := http.Client{}

	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)

		var test interface{}
		json.NewDecoder(resp.Body).Decode(&test)
		log.Println(test)

		return "", err
	}

	jsonResponse := shortenResponse{}
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	log.Printf("shortened url '%s'", targetUrl)

	return jsonResponse.Url, nil
}
