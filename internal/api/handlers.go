package api

import (
	"Echo_Dummy/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) ShowPlayers(c echo.Context) error {

	ctx := c.Request().Context()

	players, err := a.serv.ShowPlayers(ctx)

	if err != nil {
		if err == service.ErrReadingCSV {
			return c.JSON(http.StatusNotFound, errors.New("error reading csv file"))
		}
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusOK, players)

}
