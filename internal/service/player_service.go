package service

import (
	"Echo_Dummy/internal/models"
	"context"
	"errors"
)

var (
	ErrReadingCSV = errors.New("an error ocurred while reading the CSV")
)

func (s *serv) ShowPlayers(ctx context.Context) ([]models.Player, error) {

	records, err := s.repo.GetPlayers(ctx)
	if err != nil {
		return nil, ErrReadingCSV
	}

	// Player slices
	players := []models.Player{}

	// Convert the CSV records into Players
	for _, record := range records {

		// If the csv have headers with this you skip them after you save them for future purposes, if necesary
		// if i == 0 {
		// 	continue
		// }

		new_player := models.Player{
			PlayerId: record[0],
			Player:   record[1],
			Country:  record[2],
			Currency: record[3],
		}

		players = append(players, new_player)
	}

	return players, nil

}
