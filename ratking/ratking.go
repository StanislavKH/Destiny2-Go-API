package ratking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version            = "0.1"
	defaulBungieURL    = "https://www.bungie.net/"
	defaultPlatformURL = defaulBungieURL + "Platform/"
	defaultUserAgent   = "ratking/" + version
)

type ResponseData map[string]interface{}

//Possible Membership types collection
var BungieMembershipType = map[string]int{
	"All":           -1,
	"None":          0,
	"TigerXbox":     1,
	"TigerPsn":      2,
	"TigerBlizzard": 3,
	"TigerDemon":    10,
	"BungieNext":    254,
}

type Client struct {
	client      *http.Client
	BungieURL   *url.URL
	PlatformURL *url.URL
	UserAgent   string
	APIKey      string
	Character   *CharacterService
	Player      *PlayerService
	Platform    *PlatformService
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	bungieURL, _ := url.Parse(defaulBungieURL)
	platURL, _ := url.Parse(defaultPlatformURL)

	c := &Client{
		client:      httpClient,
		BungieURL:   bungieURL,
		PlatformURL: platURL,
		UserAgent:   defaultUserAgent,
		APIKey:      apiKey,
	}
	c.Platform = &PlatformService{client: c}
	c.Character = &CharacterService{client: c}
	c.Player = &PlayerService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	u := c.PlatformURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	req.Header.Add("X-API-Key", c.APIKey)
	return req, nil
}

func (cli *Client) Do(req *http.Request, v interface{}) ([]byte, error) {
	var body []byte

	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return body, err
	}
	return body, nil
}
