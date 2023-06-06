package challonge

import (
	"encoding/json"
	"fmt"
)

type TournamentInstance struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	ChallongeID uint64 `json:"id"`
	State       string `json:"state"`
}

type Tournament struct {
	Tournament TournamentInstance `json:"tournament"`
}

// Challonge docs: https://api.challonge.com/v1/documents/tournaments/index
func (c *Challonge) TournamentIndex() ([]Tournament, error) {
	params := make(map[string]string) // [TODO] Not implemented additional params
	tournaments, err := c.apiClient.TournamentIndex(params)
	if err != nil {
		return nil, err
	}

	var ret []Tournament
	json.Unmarshal(tournaments, &ret)

	for i := 0; i < len(ret); i++ {
		fmt.Printf("Tournament name: %v, URL: %v, ID: %d, State: %v\n", ret[i].Tournament.Name, ret[i].Tournament.URL, ret[i].Tournament.ChallongeID, ret[i].Tournament.State)
	}

	return ret, nil
}
