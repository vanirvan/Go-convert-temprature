package calculate

import (
	"fmt"
	"strconv"
)

func Temprature(fromUnit string, toUnit string, value float64) (float64, error) {
	const (
		fahrenheitToCelsiusFactor = 5.0 / 9.0
		celsiusToFahrenheitFactor = 9.0 / 5.0
		kelvinOffset              = 273.15
		fahrenheitKelvinOffset    = 459.67
	)

	switch {
	case fromUnit == "Fahrenheit" && toUnit == "Celsius":
		return fahrenheitToCelsiusFactor * (value - 32), nil
	case fromUnit == "Celsius" && toUnit == "Fahrenheit":
		return celsiusToFahrenheitFactor*value + 32, nil
	case fromUnit == "Kelvin" && toUnit == "Fahrenheit":
		return celsiusToFahrenheitFactor*value - fahrenheitKelvinOffset, nil
	case fromUnit == "Fahrenheit" && toUnit == "Kelvin":
		return fahrenheitToCelsiusFactor * (value + fahrenheitKelvinOffset), nil
	case fromUnit == "Kelvin" && toUnit == "Celsius":
		return value - kelvinOffset, nil
	case fromUnit == "Celsius" && toUnit == "Kelvin":
		return value + kelvinOffset, nil
	default:
		return 0, fmt.Errorf("unsupported conversion from %s to %s", fromUnit, toUnit)
	}
}

func ConvertToFloat64(value string) (float64, error) {
	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, fmt.Errorf("invalid amount, expect number instead of : %s", value)
	}

	return floatValue, nil
}
