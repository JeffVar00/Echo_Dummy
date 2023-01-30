package main

import (
	"Echo_Dummy/internal/api"
	"Echo_Dummy/internal/repository"
	"Echo_Dummy/internal/service"
	"Echo_Dummy/settings"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// Funcionalidades:
// •    Leer argumentos en JSON
// •    Leer un archivo CSV.
// •    Responder en formato JSON

// Ejemplo:
// Request:
// URL: http://localhost/getPlayers

func main() {

	app := fx.New(

		// We send all the functions that return a stroke (error)

		fx.Provide(
			context.Background,
			//database.New,
			settings.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),

		//

	)

	app.Run()

}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)

			// Start the server
			go a.Start(e, address)

			return nil

		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}
