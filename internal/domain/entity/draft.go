package entity

type PlayersDraft struct {
	PlayerName        string `csv:"player_name" json:"player_name"`
	Season            int    `csv:"season" json:"season"`
	RoundNumber       int    `csv:"round_number" json:"round_number"`
	RoundPick         int    `csv:"round_pick" json:"round_pick"`
	OverallPick       int    `csv:"overall_pick" json:"overall_pick"`
	DraftType         string `csv:"draft_type" json:"draft_type"`
	TeamID            int    `csv:"team_id" json:"team_id"`
	TeamCity          string `csv:"team_city" json:"team_city"`
	TeamName          string `csv:"team_name" json:"team_name"`
	TeamAbbreviation  string `csv:"team_abbreviation" json:"team_abbreviation"`
	Organization      string `csv:"organization" json:"organization"`
	OrganizationType  string `csv:"organization_type" json:"organization_type"`
	PlayerProfileFlag int    `csv:"player_profile_flag" json:"player_profile_flag"`
}
