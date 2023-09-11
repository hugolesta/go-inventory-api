package main

import (
	"context"

	"github.com/hugolesta/go-inventory-api/database"
	"github.com/hugolesta/go-inventory-api/internal/repository"
	"github.com/hugolesta/go-inventory-api/internal/service"
	"github.com/hugolesta/go-inventory-api/internal/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
			// func(ctx context.Context, serv service.Service){
			// 	err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	if u.Name != "myname" {
			// 		panic("invalid name")
			// 	}
				
			// },
		),
	)
	app.Run()
}
