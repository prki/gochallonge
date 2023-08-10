package challongeapi

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	BASE_URL         = "https://api.challonge.com/v1"
	TOURNAMENT_INDEX = "/tournaments.json"
	TOURNAMENT_SHOW  = "/tournaments/{tournament}.json"
)

type ChallongeAPIClient struct {
	apiKey    string
	Subdomain *string
}

func NewChallongeAPIClient(apiKey string) *ChallongeAPIClient {
	ret := &ChallongeAPIClient{
		apiKey: apiKey,
	}

	return ret
}

func (a *ChallongeAPIClient) buildURLParams(params map[string]string) url.Values {
	q := url.Values{}

	for k, v := range params {
		q.Add(k, v)
	}

	return q
}

func (a *ChallongeAPIClient) buildEndpointURL(endpoint string, params map[string]string) (*url.URL, error) {
	base, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("Error parsing API endpoint URL: %v", err)
		return nil, err
	}

	urlParams := a.buildURLParams(params)
	base.RawQuery = urlParams.Encode()

	return base, nil
}

// Method expects URL to have params encoded, if any are present
func (a *ChallongeAPIClient) get(url *url.URL) ([]byte, error) {
	response, err := http.Get(url.String())
	if err != nil {
		log.Printf("Error executing GET request: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error during IO operation reading response body during GET request: %v", err)
		return nil, err
	}

	return responseData, nil
}

func (a *ChallongeAPIClient) GetRequest(endpointURL string, params map[string]string) ([]byte, error) {
	url, err := a.buildEndpointURL(endpointURL, params)
	if err != nil {
		return nil, err
	}

	response, err := a.get(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}
