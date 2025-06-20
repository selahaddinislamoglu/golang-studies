package service

type LengthUnit string
type WeightUnit string
type TemperatureUnit string

const (
	// Length units
	Meters     LengthUnit = "meters"
	Kilometers LengthUnit = "kilometers"
	Miles      LengthUnit = "miles"
	Yards      LengthUnit = "yards"
	Feet       LengthUnit = "feet"

	// Weight units
	Kilograms WeightUnit = "kilograms"
	Pounds    WeightUnit = "pounds"
	Ounces    WeightUnit = "ounces"

	// Temperature units
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
	Kelvin     TemperatureUnit = "kelvin"
)

type Service interface {
	ConvertLength(value float64, fromUnit LengthUnit, toUnit LengthUnit) (float64, error)
	ConvertWeight(value float64, fromUnit WeightUnit, toUnit WeightUnit) (float64, error)
	ConvertTemperature(value float64, fromUnit TemperatureUnit, toUnit TemperatureUnit) (float64, error)
}
