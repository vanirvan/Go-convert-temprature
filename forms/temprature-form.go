package forms

import (
	"fmt"
	"log"

	"measurement/calculate"
	list "measurement/lists"

	"github.com/charmbracelet/huh"
)

func TempratureForm() (string, error) {
	var (
		temprature1 string
		temprature2 string
		rawAmount   string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Value(&temprature1).
				Title("Select Temprature").
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(list.Tempratures...)
				}, nil),

			huh.NewSelect[string]().
				Value(&temprature2).
				Title("Convert into").
				OptionsFunc(func() []huh.Option[string] {
					var filterTemprature []string
					for _, temprature := range list.Tempratures {
						if temprature != temprature1 {
							filterTemprature = append(filterTemprature, temprature)
						}
					}
					return huh.NewOptions(filterTemprature...)
				}, &temprature1),

			huh.NewInput().
				Value(&rawAmount).
				Title("Amount to convert"),
		).WithHeight(5),
	).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	amount, amountErr := calculate.ConvertToFloat64(rawAmount)

	if amountErr != nil {
		return "", amountErr
	}

	result, resultError := calculate.Temprature(temprature1, temprature2, float64(amount))

	if resultError != nil {
		return "", resultError
	}

	return fmt.Sprintf("%.2f° %s to %s is %.2f°", float64(amount), temprature1, temprature2, result), nil

}
