package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/service"
)

type UnitConverterController struct {
	service service.Service
}

func NewUnitConverterController(service service.Service) Controller {
	return &UnitConverterController{
		service: service,
	}
}

var templates = template.Must(template.ParseGlob("internal/templates/*.html"))

func (c *UnitConverterController) Home(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		data := map[string]interface{}{
			"Title":   "Unit Converter",
			"Content": "home",
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	} else {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (c *UnitConverterController) Length(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		data := map[string]interface{}{
			"Title":   "Length Converter",
			"Content": "convert",
			"Unit":    "length",
			"Units":   []string{"meters", "kilometers", "miles", "yards", "feet"},
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	case http.MethodPost:
		value := request.FormValue("value")
		from := request.FormValue("from")
		to := request.FormValue("to")

		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(response, "Invalid input value", http.StatusBadRequest)
			return
		}

		result, err := c.service.ConvertLength(valueFloat, service.LengthUnit(from), service.LengthUnit(to))
		if err != nil {
			http.Error(response, "Conversion error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":          "Length Converter",
			"Content":        "result",
			"Unit":           "length",
			"Value":          value,
			"From":           from,
			"ConvertedValue": result,
			"To":             to,
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (c *UnitConverterController) Weight(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		data := map[string]interface{}{
			"Title":   "Weight Converter",
			"Content": "convert",
			"Unit":    "weight",
			"Units":   []string{"kilograms", "pounds", "ounces"},
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	case http.MethodPost:
		value := request.FormValue("value")
		from := request.FormValue("from")
		to := request.FormValue("to")

		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(response, "Invalid input value", http.StatusBadRequest)
			return
		}

		result, err := c.service.ConvertWeight(valueFloat, service.WeightUnit(from), service.WeightUnit(to))
		if err != nil {
			http.Error(response, "Conversion error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":          "Weight Converter",
			"Content":        "result",
			"Unit":           "weight",
			"Value":          value,
			"From":           from,
			"ConvertedValue": result,
			"To":             to,
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (c *UnitConverterController) Temperature(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		data := map[string]interface{}{
			"Title":   "Temperature Converter",
			"Content": "convert",
			"Unit":    "temperature",
			"Units":   []string{"celsius", "fahrenheit", "kelvin"},
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	case http.MethodPost:
		value := request.FormValue("value")
		from := request.FormValue("from")
		to := request.FormValue("to")

		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(response, "Invalid input value", http.StatusBadRequest)
			return
		}

		result, err := c.service.ConvertTemperature(valueFloat, service.TemperatureUnit(from), service.TemperatureUnit(to))
		if err != nil {
			http.Error(response, "Conversion error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":          "Temperature Converter",
			"Content":        "result",
			"Unit":           "temperature",
			"Value":          value,
			"From":           from,
			"ConvertedValue": result,
			"To":             to,
		}
		templates.ExecuteTemplate(response, "layout.html", data)
	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
