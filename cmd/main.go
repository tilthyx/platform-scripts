package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tilthyx/platform-scripts/cmd/csv"
	"github.com/tilthyx/platform-scripts/cmd/database"
	"github.com/tilthyx/platform-scripts/internal/domain/entity"
)

// Map of structs keyed by file name
var fileStructMap = map[string]interface{}{
	"teams.csv":   entity.Teams{},
	"players.csv": entity.Players{},
}

func main() {
	// Reading all files in assets
	files, err := csv.GetCSVFileNames("./assets")
	if err != nil {
		log.Fatal(err)
	}

	var header []string
	var lines []interface{}

	// Create db connection
	db, err := database.Connection()
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			_ = fmt.Errorf("unable to close the database: %v", err)
		}
	}(db)

	// Reader the csv file
	for _, file := range files {
		header, lines, err = csv.ReaderCSV(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Header: %v\n", header)
		_, err := json.MarshalIndent(lines, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		// Insert records into DB
		for _, item := range lines {
			switch v := item.(type) {
			case *entity.Teams:
				// Insert Team data
				err := database.InsertTeam(db, *v)
				if err != nil {
					fmt.Printf("Error inserting team into DB: %v\n", err)
				} else {
					fmt.Printf("Successfully inserted team: %s\n", v.Name)
				}
			case *entity.Players:
				// Insert Player data
				err := database.InsertPlayer(db, *v)
				if err != nil {
					fmt.Printf("Error inserting player into DB: %v\n", err)
				} else {
					fmt.Printf("Successfully inserted player: %s\n", v.PlayerName)
				}
			default:
				fmt.Println("Unknown type for insertion")
			}
		}
	}

}
