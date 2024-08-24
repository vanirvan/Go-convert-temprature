package main

import (
	"fmt"
	"measurement/forms"
)

func main() {

	var (
		result      any
		errorResult error
	)

	var initialFormValue = forms.InitialForm()
	switch initialFormValue {
	case "Temprature":
		result, errorResult = forms.TempratureForm()
		if errorResult != nil {
			fmt.Println(errorResult)
			break
		}

		fmt.Println(result)
	default:
		fmt.Println("Program closed")
	}

	// * Loop, do it later
	// isLoop := true

	// for {

	// 	if !isLoop {
	// 		fmt.Println("See you again!")
	// 		break
	// 	}
	// }
}
