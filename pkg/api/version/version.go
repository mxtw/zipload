package version

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/mxtw/zipload/pkg/api"
)

type versionInfo struct {
	Version string `json:"version"`
}

func Version(client *api.Client) (versionInfo, error) {
	endpoint, err := url.JoinPath(client.Host, "/api/version")
	if err != nil {
		log.Println(err)
		return versionInfo{}, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return versionInfo{}, err
	}

	req.Header.Add("Authorization", client.Token)

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err)
		return versionInfo{}, err
	}

	jsonResponse := versionInfo{}
	json.NewDecoder(resp.Body).Decode(&jsonResponse)
	defer resp.Body.Close()

	return jsonResponse, nil
}
