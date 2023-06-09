package challongeapi

import (
	"log"
	"strings"
)

func (a *ChallongeAPIClient) TournamentIndex(params map[string]string) ([]byte, error) {
	endpointURL := BASE_URL + TOURNAMENT_INDEX
	params["api_key"] = a.apiKey

	response, err := a.GetRequest(endpointURL, params)
	if err != nil {
		log.Printf("Error executing GET request for Tournament Index: %v", err)
	}

	return response, nil
}

func (a *ChallongeAPIClient) TournamentShow(tournamentID string, params map[string]string) ([]byte, error) {
	endpointURL := BASE_URL + TOURNAMENT_SHOW
	endpointURL = strings.Replace(endpointURL, "{tournament}", tournamentID, 1)

	params["api_key"] = a.apiKey

	response, err := a.GetRequest(endpointURL, params)
	if err != nil {
		log.Printf("Error executing GET request for Tournament Show: %v", err)
		return nil, err
	}

	return response, nil
}
