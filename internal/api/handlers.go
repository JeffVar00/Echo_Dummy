package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) ShowPlayers(c echo.Context) error {

	ctx := c.Request().Context()

	players, err := a.serv.ShowPlayers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusOK, players)
}
