package repository

import (
	//"Echo_Dummy/internal/entity"
	"context"
)

// Repository is the interface that wraps the basic CRUD methods

//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	GetPlayers(ctx context.Context, filename string) ([][]string, error)
}

// Insted of using a [][]string, we should use an entity, but since we are not working with a database, it is not necessary

// Receive the interfaces and report structs

// Database Interface for Repository
type repo struct {
	//db *sqlx.DB
}

func New() Repository {
	return &repo{}
}

// Init Repository
// func New(db *sqlx.DB) Repository {
// 	return &repo{db: db}
// }

//	Install and use Mockery
//	go install github.com/vektra/mockery/v2@latest
//	go generate ./... //esto lo hace de forma recursiva
