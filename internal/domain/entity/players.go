package entity

type Players struct {
	PlayerName       string  `csv:"player_name" json:"player_name"`             // Player's name
	TeamAbbreviation string  `csv:"team_abbreviation" json:"team_abbreviation"` // Team abbreviation
	Age              int     `csv:"age" json:"age"`                             // Age
	PlayerHeight     float64 `csv:"player_height" json:"player_height"`         // Height in cm
	PlayerWeight     float64 `csv:"player_weight" json:"player_weight"`         // Weight in kg
	College          string  `csv:"college" json:"college"`                     // College name
	Country          string  `csv:"country" json:"country"`                     // Country
	DraftYear        int     `csv:"draft_year" json:"draft_year"`               // Draft year
	DraftRound       int     `csv:"draft_round" json:"draft_round"`             // Draft round
	DraftNumber      int     `csv:"draft_number" json:"draft_number"`           // Draft number
	GP               int     `csv:"gp" json:"gp"`                               // Games played
	PTS              float64 `csv:"pts" json:"pts"`                             // Points per game
	REB              float64 `csv:"reb" json:"reb"`                             // Rebounds per game
	AST              float64 `csv:"ast" json:"ast"`                             // Assists per game
	NetRating        float64 `csv:"net_rating" json:"net_rating"`               // Net rating
	OREBPercent      float64 `csv:"oreb_pct" json:"oreb_pct"`                   // Offensive rebound percentage
	DREBPercent      float64 `csv:"dreb_pct" json:"dreb_pct"`                   // Defensive rebound percentage
	USGPercent       float64 `csv:"usg_pct" json:"usg_pct"`                     // Usage percentage
	TSPercent        float64 `csv:"ts_pct" json:"ts_pct"`                       // True shooting percentage
	ASTPercent       float64 `csv:"ast_pct" json:"ast_pct"`                     // Assist percentage
	Season           string  `csv:"season" json:"season"`                       // Season
}
