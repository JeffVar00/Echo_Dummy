package service

import (
	"Echo_Dummy/internal/models"
	"Echo_Dummy/internal/repository"
	"context"
)

// Service is the bussiness logic of the application

//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	// Return the list of players from the repository in a JSON format
	ShowPlayers(ctx context.Context) ([]models.Player, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
