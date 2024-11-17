package entity

type TeamsDetails struct {
	Abbreviation       string `json:"abbreviation" csv:"abbreviation"`
	Nickname           string `json:"nickname" csv:"nickname"`
	YearFounded        int    `json:"yearfounded" csv:"yearfounded"`
	City               string `json:"city" csv:"city"`
	Arena              string `json:"arena" csv:"arena"`
	ArenaCapacity      int    `json:"arenacapacity" csv:"arenacapacity"`
	Owner              string `json:"owner" csv:"owner"`
	GeneralManager     string `json:"generalmanager" csv:"generalmanager"`
	HeadCoach          string `json:"headcoach" csv:"headcoach"`
	DLeagueAffiliation string `json:"dleagueaffiliation" csv:"dleagueaffiliation"`
	Facebook           string `json:"facebook" csv:"facebook"`
	Instagram          string `json:"instagram" csv:"instagram"`
	Twitter            string `json:"twitter" csv:"twitter"`
}
