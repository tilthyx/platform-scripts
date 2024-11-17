package database

import (
	"database/sql"
	"fmt"
	"github.com/tilthyx/platform-scripts/internal/domain/entity"
)

func InsertTeam(db *sql.DB, team entity.Teams) error {
	// Define the SQL Insert query with placeholders
	query := `INSERT INTO teams (name, short_name, slug) VALUES ($1, $2, $3)`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, team.Name, team.ShortName, team.Slug)
	if err != nil {
		return fmt.Errorf("could not insert team: %v", err)
	}

	return nil
}

func InsertPlayer(db *sql.DB, player entity.Players) error {
	// Define the SQL Insert query with placeholders
	query := `INSERT INTO players (
		player_name, team_abbreviation, age, player_height, player_weight, college, country, draft_year, draft_round,
		draft_number, gp, pts, reb, ast, net_rating, oreb_pct, dreb_pct, usg_pct, ts_pct, ast_pct, season
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, player.PlayerName, player.TeamAbbreviation, player.Age, player.PlayerHeight, player.PlayerWeight,
		player.College, player.Country, player.DraftYear, player.DraftRound, player.DraftNumber, player.GP,
		player.PTS, player.REB, player.AST, player.NetRating, player.OREBPercent, player.DREBPercent, player.USGPercent,
		player.TSPercent, player.ASTPercent, player.Season)

	if err != nil {
		return fmt.Errorf("could not insert player: %v", err)
	}

	return nil
}
