package repository

import (
	//"Echo_Dummy/internal/entity"
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

const (
// QUERYS
)

func (r *repo) GetPlayers(ctx context.Context, filename string) ([][]string, error) {

	// LOGIC CSV READ FILE // SIMULATION OF A DATABASE EXTRACTION
	// Open the CSV file

	file := fmt.Sprintf("internal/data/%s.csv", filename)

	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	// Read the CSV file
	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
