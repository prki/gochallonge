package challonge

import (
	"log"

	"github.com/prki/gochallonge/pkg/challongeapi"
)

type Challonge struct {
	apiClient *challongeapi.ChallongeAPIClient
}

func NewChallonge(apiKey string) *Challonge {
	if apiKey == "" {
		log.Print("Unable to initialize Challonge connection with empty API key")
		return nil
	}
	apiClient := challongeapi.NewChallongeAPIClient(apiKey)

	ret := &Challonge{
		apiClient: apiClient,
	}

	return ret
}
