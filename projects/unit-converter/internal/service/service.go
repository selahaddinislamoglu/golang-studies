package service

type UnitConverterService struct{}

func NewUnitConverterService() Service {
	return &UnitConverterService{}
}

func (s *UnitConverterService) ConvertLength(value float64, fromUnit LengthUnit, toUnit LengthUnit) (float64, error) {

	switch fromUnit {
	case Meters:
		switch toUnit {
		case Meters:
			return value, nil
		case Kilometers:
			return value / 1000, nil
		case Miles:
			return value / 1609.34, nil
		case Yards:
			return value * 1.09361, nil
		case Feet:
			return value * 3.28084, nil
		default:
			return 0.0, nil
		}
	case Kilometers:
		switch toUnit {
		case Meters:
			return value * 1000, nil
		case Kilometers:
			return value, nil
		case Miles:
			return value / 1.60934, nil
		case Yards:
			return value * 1093.61, nil
		case Feet:
			return value * 3280.84, nil
		default:
			return 0.0, nil
		}
	case Miles:
		switch toUnit {
		case Meters:
			return value * 1609.34, nil
		case Kilometers:
			return value * 1.60934, nil
		case Miles:
			return value, nil
		case Yards:
			return value * 1760, nil
		case Feet:
			return value * 5280, nil
		default:
			return 0.0, nil
		}
	case Yards:
		switch toUnit {
		case Meters:
			return value / 1.09361, nil
		case Kilometers:
			return value / 1093.61, nil
		case Miles:
			return value / 1760, nil
		case Yards:
			return value, nil
		case Feet:
			return value * 3, nil
		default:
			return 0.0, nil
		}
	case Feet:
		switch toUnit {
		case Meters:
			return value / 3.28084, nil
		case Kilometers:
			return value / 3280.84, nil
		case Miles:
			return value / 5280, nil
		case Yards:
			return value / 3, nil
		case Feet:
			return value, nil
		default:
			return 0.0, nil
		}
	default:
		return 0.0, nil
	}
}

func (s *UnitConverterService) ConvertWeight(value float64, fromUnit WeightUnit, toUnit WeightUnit) (float64, error) {

	switch fromUnit {
	case Kilograms:
		switch toUnit {
		case Kilograms:
			return value, nil
		case Pounds:
			return value * 2.20462, nil
		case Ounces:
			return value * 35.274, nil
		default:
			return 0.0, nil
		}
	case Pounds:
		switch toUnit {
		case Kilograms:
			return value / 2.20462, nil
		case Pounds:
			return value, nil
		case Ounces:
			return value * 16, nil
		default:
			return 0.0, nil
		}
	case Ounces:
		switch toUnit {
		case Kilograms:
			return value / 35.274, nil
		case Pounds:
			return value / 16, nil
		case Ounces:
			return value, nil
		default:
			return 0.0, nil
		}
	default:
		return 0.0, nil
	}
}

func (s *UnitConverterService) ConvertTemperature(value float64, fromUnit TemperatureUnit, toUnit TemperatureUnit) (float64, error) {

	switch fromUnit {
	case Celsius:
		switch toUnit {
		case Celsius:
			return value, nil
		case Fahrenheit:
			return value*1.8 + 32, nil
		case Kelvin:
			return value + 273.15, nil
		default:
			return 0.0, nil
		}
	case Fahrenheit:
		switch toUnit {
		case Celsius:
			return (value - 32) / 1.8, nil
		case Fahrenheit:
			return value, nil
		case Kelvin:
			return (value-32)/1.8 + 273.15, nil
		default:
			return 0.0, nil
		}
	case Kelvin:
		switch toUnit {
		case Celsius:
			return value - 273.15, nil
		case Fahrenheit:
			return (value-273.15)*1.8 + 32, nil
		case Kelvin:
			return value, nil
		default:
			return 0.0, nil
		}
	default:
		return 0.0, nil
	}
}
