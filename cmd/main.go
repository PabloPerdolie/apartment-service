package main

import (
	"apartment_search_service/internal/app"
	"context"
)

func main() {
	a := app.NewApp(context.Background())

	a.Run()
}
