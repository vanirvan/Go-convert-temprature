package main

import (
	"fmt"
	"measurement/forms"

	"github.com/charmbracelet/huh"
)

func main() {
	isLoop := true

	for {
		var again bool = true

		result, error := forms.TempratureForm()

		if error != nil {
			fmt.Println(error)
			break
		}

		fmt.Println(result)

		confirmForm := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Want to do it again?").
					Value(&again),
			),
		).WithTheme(huh.ThemeCatppuccin())

		confirmForm.Run()

		if !isLoop || !again {
			fmt.Println("See you again!")
			break
		}
	}
}
