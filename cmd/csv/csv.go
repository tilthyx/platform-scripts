package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/tilthyx/platform-scripts/internal/domain/entity"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

// Map of structs keyed by file name
var fileStructMap = map[string]interface{}{
	"teams.csv":   &entity.Teams{},   // Pointers to structs
	"players.csv": &entity.Players{}, // Pointers to structs
}

func ReaderCSV(csvName string) ([]string, []interface{}, error) {
	// Open the CSV file
	fileName := "./assets/" + csvName
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, nil, fmt.Errorf("CSV file must have at least a header and one data row")
	}

	header := records[0] // Extract the header row
	var result []interface{}

	// Use reflection to process each record into the correct struct type
	for _, record := range records[1:] {
		var structInstance interface{}
		switch csvName {
		case "teams.csv":
			// Create a new Team struct instance
			structInstance = &entity.Teams{} // Use pointer for reflection
		case "players.csv":
			// Create a new Player struct instance
			structInstance = &entity.Players{} // Use pointer for reflection
		default:
			return nil, nil, fmt.Errorf("no struct defined for file: %s", csvName)
		}

		// Use reflection to map fields
		structValue := reflect.ValueOf(structInstance).Elem()

		for i, fieldName := range header {
			// Use reflection to map fields via struct tags (`csv` tags)
			// Find the struct field by `csv` tag
			structField := getStructFieldByCSVTag(structValue, fieldName)

			// If the field exists and is valid, set the field value
			if structField != "" {
				field := structValue.FieldByName(structField)
				if field.IsValid() && field.CanSet() {
					// Set the field value
					setFieldValue(field, record[i])
				}
			}
		}
		result = append(result, structInstance)
	}

	return header, result, nil
}

// getStructFieldByCSVTag finds the struct field name based on the `csv` tag
func getStructFieldByCSVTag(v reflect.Value, tag string) string {
	for i := 0; i < v.Type().NumField(); i++ {
		field := v.Type().Field(i)
		if csvTag := field.Tag.Get("csv"); csvTag == tag {
			return field.Name
		}
	}
	return ""
}

// setFieldValue is a helper to set the correct field based on type
func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		if intValue, err := strconv.Atoi(value); err == nil {
			field.SetInt(int64(intValue))
		}
	case reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			field.SetFloat(floatValue)
		}
		// Add other types as needed
	default:
		fmt.Println("Unhandled type for field:", field.Kind())
	}
}

func GetCSVFileNames(folderPath string) ([]string, error) {
	var csvFileNames []string

	// Walk through the directory
	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .csv extension
		if !d.IsDir() && filepath.Ext(d.Name()) == ".csv" {
			csvFileNames = append(csvFileNames, filepath.Base(path)) // Add only the file name
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return csvFileNames, nil
}
