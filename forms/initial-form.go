package forms

import (
	"log"

	"github.com/charmbracelet/huh"
)

func InitialForm() string {
	var measureResult string
	var measureList = []string{
		"Temprature",
		"Weight",
		"Height",
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What do you want to measure").
				Value(&measureResult).
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(measureList...)
				}, nil).
				Height(5),
		),
	).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	return measureResult
}
