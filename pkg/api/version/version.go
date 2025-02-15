package version

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/mxtw/zipload/pkg/api"
)

type versions struct {
	Stable   string `json:"stable"`
	Upstream string `json:"upstream"`
	Current  string `json:"current"`
}

type versionInfo struct {
	IsUpstream   bool     `json:"isUpstream"`
	Update       bool     `json:"update"`
	UpdateToType string   `json:"updateToType"`
	Versions     versions `json:"versions"`
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
