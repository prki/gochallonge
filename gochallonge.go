package gochallonge

import "github.com/prki/gochallonge/pkg/challonge"

// Access function for public usage of library
func New(apiKey string) *challonge.Challonge {
	return challonge.NewChallonge(apiKey)
}
