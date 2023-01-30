package service

import (
	"Echo_Dummy/internal/repository"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {

	repo = &repository.MockRepository{}

	// Extract the data correctly from the database, in this case from the CSV
	repo.On("GetPlayers", mock.Anything).Return(nil, nil)

	s = New(repo)

	code := m.Run() //Return a Status Code
	os.Exit(code)   //Allows to exit the program with a status code

}

// UNIT TESTS

// Normally I only the bussisnes logic, but for this case I have to test the conection
func TestShowPlayers(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "Show_Players_Success",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {

			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.ShowPlayers(ctx)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}

		})
	}

}
