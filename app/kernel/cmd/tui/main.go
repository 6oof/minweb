package main

import (
	"fmt"
	"log"

	"github.com/6oof/minweb/app/kernel/cmd/tui/migrations"
	"github.com/6oof/minweb/app/kernel/cmd/tui/xxhtml"
	"github.com/charmbracelet/huh"
)

var (
	op int
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[int]().
				Title("Select operation:").
				Options(
					huh.NewOption("Transform HTML to XXHTML", 1),
					huh.NewOption("Migrations", 2),
				).
				Value(&op), // store the chosen option in the "burger" variable
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch o := op; o {
	case 1:
		xxhtml.Run()
	case 2:
		migrations.Run()
	default:
		fmt.Println("No option selected")
	}
}
