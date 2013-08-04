package dota2

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	SteamAPIHost = "api.steampowered.com"
)

type Dota2Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewDota2Client(apiKey string) *Dota2Client {
	return &Dota2Client{APIKey: apiKey, HTTPClient: &http.Client{}}
}

func (d *Dota2Client) GetLeagueListing() []*League {
	v := url.Values{}
	v.Set("key", d.APIKey)
	url := url.URL{
		Scheme:   "https",
		Host:     SteamAPIHost,
		Path:     "/IDOTA2Match_570/GetLeagueListing/v001",
		RawQuery: v.Encode(),
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		panic(err)
	}
	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return UnmarshalLeagues(b)
}
