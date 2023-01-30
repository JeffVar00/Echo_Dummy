package main

import (
	"Echo_Dummy/internal/repository"
	"Echo_Dummy/internal/service"
	"Echo_Dummy/settings"
	"context"

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
			settings.New,
			repository.New,
			service.New,
		),
		fx.Invoke(

			// direct testing

			func(ctx context.Context, serv service.Service) {
				_, err := serv.ShowPlayers(ctx)
				if err != nil {
					panic(err)
				}
			},

			//

		),
	)
	app.Run()
}
