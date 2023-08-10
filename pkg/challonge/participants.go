package challonge

type Participant struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	FinalRank         int    `json:"final_rank"`
	ChallongeUsername string `json:"challonge_username"`
	ChallongeUserId   int    `json:"challonge_user_id"`
	DisplayName       string `json:"display_name"`
}

type ParticipantResponse struct {
	Participant *Participant `json:"participant"`
}
