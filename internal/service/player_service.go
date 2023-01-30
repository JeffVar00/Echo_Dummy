package service

import (
	"Echo_Dummy/internal/models"
	"context"
	"errors"
)

var (
	ErrCreatingJSON = errors.New("an error ocurred while creating the JSON")
	ErrReadingCSV   = errors.New("an error ocurred while reading the CSV")
)

func (s *serv) ShowPlayers(ctx context.Context) ([]models.Player, error) {

	records, err := s.repo.GetPlayers(ctx)
	if err != nil {
		return nil, ErrReadingCSV
	}

	// Player slices
	players := []models.Player{}

	// Convert the CSV records into Players
	for i, record := range records {
		if i == 0 {
			continue
		}

		new_player := models.Player{
			PlayerId: record[0],
			Player:   record[1],
			Country:  record[2],
			Currency: record[3],
		}

		players = append(players, new_player)
	}

	// THIS IS DONE IN THE API LAYER
	// Convert the data to a JSON string
	// playersData, err := json.Marshal(players)
	// if err != nil {
	// 	return nil, ErrCreatingJSON
	// }

	return players, nil

}
