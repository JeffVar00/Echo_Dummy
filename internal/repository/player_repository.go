package repository

import (
	//"Echo_Dummy/internal/entity"
	"context"
	"encoding/csv"
	"os"
)

const (
// QUERYS
)

func (r *repo) GetPlayers(ctx context.Context) ([][]string, error) {

	//LOGICA CSV READ FILE // SIMULATION OF A DATABASE EXTRACTION

	// Open the CSV file
	csvFile, err := os.Open("C:/Users/jeff0/Documents/GRW Analytics/Echo_Dummy/internal/data/data.csv")
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
