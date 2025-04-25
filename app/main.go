package main

import (
	"front-exercise/presentation"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		presentation.PresentationModule,
	).Run()
}
