package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {
	players := e.Group("/players")
	players.POST("", a.ShowPlayers)
}
