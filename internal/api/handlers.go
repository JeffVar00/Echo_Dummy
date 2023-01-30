package api

import (
	"Echo_Dummy/internal/api/dtos"
	"Echo_Dummy/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BadRequestMessage struct {
	Message string `json:"message"`
}

func (a *API) ShowPlayers(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.SearchFile{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequestMessage{Message: "Invalid Request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequestMessage{Message: err.Error()})
	}

	players, err := a.serv.ShowPlayers(ctx, params.FileName)

	if err != nil {
		if err == service.ErrCSVNotFound {
			return c.JSON(http.StatusNotFound, errors.New("error reading csv file"))
		}
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusOK, players)

}
