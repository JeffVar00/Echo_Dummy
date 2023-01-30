package api

import (
	"Echo_Dummy/internal/service"

	"github.com/labstack/echo/v4"
)

// VALIDATORS
// GO GET github.com/go-playground/validator/v10

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{serv}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}
