package challongeapi

func (a *ChallongeAPIClient) TournamentIndex(params map[string]string) ([]byte, error) {
	endpointURL := BASE_URL + TOURNAMENT_INDEX
	params["api_key"] = a.apiKey

	url, err := a.buildEndpointURL(endpointURL, params)
	if err != nil {
		return nil, err
	}

	response, err := a.GetRequest(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}
