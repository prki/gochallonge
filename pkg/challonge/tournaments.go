package challonge

import (
	"encoding/json"
)

type Tournament struct {
	Name         string                `json:"name"`
	URL          string                `json:"url"`
	ChallongeID  uint64                `json:"id"`
	State        string                `json:"state"`
	StartAt      string                `json:"start_at"`
	StartedAt    string                `json:"started_at"`
	Participants []ParticipantResponse `json:"participants,omitempty"`
}

type TournamentResponse struct {
	Tournament Tournament `json:"tournament"`
}

// Challonge docs: https://api.challonge.com/v1/documents/tournaments/index
func (c *Challonge) TournamentIndex() ([]TournamentResponse, error) {
	params := make(map[string]string) // [TODO] Not implemented additional params
	tournaments, err := c.apiClient.TournamentIndex(params)
	if err != nil {
		return nil, err
	}

	var ret []TournamentResponse
	json.Unmarshal(tournaments, &ret)

	return ret, nil
}

func (c *Challonge) TournamentShow(tournamentID string, includeParticipants bool, includeMatches bool) (*TournamentResponse, error) {
	params := make(map[string]string)
	if includeParticipants {
		params["include_participants"] = "1"
	}
	if includeMatches {
		params["include_matches"] = "1"
	}

	tournament, err := c.apiClient.TournamentShow(tournamentID, params)
	if err != nil {
		return nil, err
	}

	ret := new(TournamentResponse)
	json.Unmarshal(tournament, ret)

	return ret, nil
}
