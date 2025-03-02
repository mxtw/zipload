package urls

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mxtw/zipload/pkg/api"
)

// TODO: centralize those options somewhere
type Options struct {
	MaxViews uint
	Domain   string
	Password string
}

type postBody struct {
	Destination string `json:"destination"`
	Vanity      string `json:"vanity"`
}

type shortenResponse struct {
	Url string `json:"url"`
}

func Url(client *api.Client, targetUrl string, vanity string, options Options) (string, error) {
	endpoint, err := url.JoinPath(client.Host, "/api/user/urls")
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

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", client.Token)

	if options.MaxViews > 0 {
		req.Header.Add("x-zipline-max-views", strconv.FormatUint(uint64(options.MaxViews), 10))
	}
	if options.Domain != "" {
		req.Header.Add("x-zipline-domain", options.Domain)
	}
	if options.Password != "" {
		req.Header.Add("x-zipline-password", options.Password)
	}

	hc := http.Client{}

	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)

		return "", err
	}

	jsonResponse := shortenResponse{}
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	log.Printf("shortened url '%s'", targetUrl)

	return jsonResponse.Url, nil
}
