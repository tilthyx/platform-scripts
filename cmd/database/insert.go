package database

import (
	"database/sql"
	"fmt"
	"github.com/tilthyx/platform-scripts/internal/domain/entity"
)

// Insert or update team if abbreviation exists
func InsertTeam(db *sql.DB, team entity.Teams) error {
	// Define the SQL Insert query with ON CONFLICT DO UPDATE
	query := `
		INSERT INTO teams (name, short_name, slug)
		VALUES ($1, $2, $3)
		ON CONFLICT (name) 
		DO UPDATE SET 
			short_name = EXCLUDED.short_name,
			slug = EXCLUDED.slug
	`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, team.Name, team.ShortName, team.Slug)
	if err != nil {
		return fmt.Errorf("could not insert or update team: %v", err)
	}

	return nil
}

// Insert or update player if player name exists
func InsertPlayer(db *sql.DB, player entity.Players) error {
	// Define the SQL Insert query with ON CONFLICT DO UPDATE
	query := `
		INSERT INTO players (
			player_name, team_abbreviation, age, player_height, player_weight, college, country, 
			draft_year, draft_round, draft_number, gp, pts, reb, ast, net_rating, 
			oreb_pct, dreb_pct, usg_pct, ts_pct, ast_pct, season
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
		ON CONFLICT (player_name) 
		DO UPDATE SET 
			team_abbreviation = EXCLUDED.team_abbreviation, 
			age = EXCLUDED.age,
			player_height = EXCLUDED.player_height,
			player_weight = EXCLUDED.player_weight,
			college = EXCLUDED.college,
			country = EXCLUDED.country,
			draft_year = EXCLUDED.draft_year,
			draft_round = EXCLUDED.draft_round,
			draft_number = EXCLUDED.draft_number,
			gp = EXCLUDED.gp,
			pts = EXCLUDED.pts,
			reb = EXCLUDED.reb,
			ast = EXCLUDED.ast,
			net_rating = EXCLUDED.net_rating,
			oreb_pct = EXCLUDED.oreb_pct,
			dreb_pct = EXCLUDED.dreb_pct,
			usg_pct = EXCLUDED.usg_pct,
			ts_pct = EXCLUDED.ts_pct,
			ast_pct = EXCLUDED.ast_pct,
			season = EXCLUDED.season
	`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, player.PlayerName, player.TeamAbbreviation, player.Age, player.PlayerHeight, player.PlayerWeight,
		player.College, player.Country, player.DraftYear, player.DraftRound, player.DraftNumber, player.GP,
		player.PTS, player.REB, player.AST, player.NetRating, player.OREBPercent, player.DREBPercent, player.USGPercent,
		player.TSPercent, player.ASTPercent, player.Season)

	if err != nil {
		return fmt.Errorf("could not insert or update player: %v", err)
	}

	return nil
}

// Insert or update players draft info if player name and season exist
func InsertPlayersDraft(db *sql.DB, player entity.PlayersDraft) error {
	// Define the SQL Insert query with ON CONFLICT DO UPDATE
	query := `
		INSERT INTO players_draft (
			player_name, season, round_number, round_pick, overall_pick, draft_type, 
			team_id, team_city, team_name, team_abbreviation, organization, 
			organization_type, player_profile_flag
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		ON CONFLICT (player_name) 
		DO UPDATE SET 
			season = EXCLUDED.season,
			round_number = EXCLUDED.round_number,
			round_pick = EXCLUDED.round_pick,
			overall_pick = EXCLUDED.overall_pick,
			draft_type = EXCLUDED.draft_type,
			team_id = EXCLUDED.team_id,
			team_city = EXCLUDED.team_city,
			team_name = EXCLUDED.team_name,
			team_abbreviation = EXCLUDED.team_abbreviation,
			organization = EXCLUDED.organization,
			organization_type = EXCLUDED.organization_type,
			player_profile_flag = EXCLUDED.player_profile_flag
	`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, player.PlayerName, player.Season, player.RoundNumber, player.RoundPick, player.OverallPick,
		player.DraftType, player.TeamID, player.TeamCity, player.TeamName, player.TeamAbbreviation,
		player.Organization, player.OrganizationType, player.PlayerProfileFlag)

	if err != nil {
		return fmt.Errorf("could not insert or update player draft: %v", err)
	}

	return nil
}

// Insert or update team details if abbreviation exists
func InsertTeamDetails(db *sql.DB, team entity.TeamsDetails) error {
	// Define the SQL Insert query with ON CONFLICT DO UPDATE
	query := `
		INSERT INTO team_details (
			abbreviation, nickname, yearfounded, city, arena, arenacapacity, 
			owner, generalmanager, headcoach, dleagueaffiliation, 
			facebook, instagram, twitter
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		ON CONFLICT (abbreviation) 
		DO UPDATE SET 
			nickname = EXCLUDED.nickname, 
			yearfounded = EXCLUDED.yearfounded, 
			city = EXCLUDED.city, 
			arena = EXCLUDED.arena, 
			arenacapacity = EXCLUDED.arenacapacity,
			owner = EXCLUDED.owner,
			generalmanager = EXCLUDED.generalmanager,
			headcoach = EXCLUDED.headcoach,
			dleagueaffiliation = EXCLUDED.dleagueaffiliation,
			facebook = EXCLUDED.facebook,
			instagram = EXCLUDED.instagram,
			twitter = EXCLUDED.twitter
	`

	// Execute the query with the struct fields as arguments
	_, err := db.Exec(query, team.Abbreviation, team.Nickname, team.YearFounded, team.City, team.Arena, team.ArenaCapacity,
		team.Owner, team.GeneralManager, team.HeadCoach, team.DLeagueAffiliation,
		team.Facebook, team.Instagram, team.Twitter)

	if err != nil {
		return fmt.Errorf("could not insert or update team details: %v", err)
	}

	return nil
}
