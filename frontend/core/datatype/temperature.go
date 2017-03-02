package datatype

import (
	"fmt"
)

type tempType int

// Temperature units - Celsius, Fahrenheit, Kelvin
const (
	Celsius tempType = iota
	Fahrenheit
	Kelvin
)

// TemperatureType - struct for implementing convertions of temperature
type TemperatureType struct {
	b             *BaseDataType
	tempUnit      tempType
	convTempFuncs map[tempType]ConvFunc
}

// GetBase - return basic info(names, displayName, unit group) of datatype unit
func (t *TemperatureType) GetBase() *BaseDataType {
	return t.b
}

// GetConvFunc - convert function for switching from one temperature unit to antoher
func (t *TemperatureType) GetConvFunc(typeTo DataType) (ConvFunc, error) {
	if t.b.Group == GroupBare {
		return func(in float64) float64 { return in }, nil
	} else if t.b.Group != typeTo.GetBase().Group {
		return nil, fmt.Errorf("GetConversionMultipl: incompatible types %#v - %#v", t, typeTo)
	}

	return t.convTempFuncs[typeTo.(*TemperatureType).tempUnit], nil
}

var temperatureTypes = []DataType{
	&TemperatureType{
		b: &BaseDataType{
			Group:       GroupTemperature,
			Names:       []string{"celsius", "grad celsius", "grads celsius", "C", "°C"},
			DisplayName: "°C",
		},
		tempUnit: Celsius,
		convTempFuncs: map[tempType]ConvFunc{
			Fahrenheit: func(in float64) float64 { return 1.8*in + 32 },
			Kelvin:     func(in float64) float64 { return in + 273.15 },
		},
	},
	&TemperatureType{
		b: &BaseDataType{
			Group:       GroupTemperature,
			Names:       []string{"fahrenheit", "grad fahrenheit", "grads fahrenheit", "F", "°F"},
			DisplayName: "°F",
		},
		tempUnit: Fahrenheit,
		convTempFuncs: map[tempType]ConvFunc{
			Celsius: func(in float64) float64 { return (in - 32) / 1.8 },
			Kelvin:  func(in float64) float64 { return (in-32)/1.8 + 273.15 },
		},
	},
	&TemperatureType{
		b: &BaseDataType{
			Group:       GroupTemperature,
			Names:       []string{"kelvin", "kelvins", "K"},
			DisplayName: "K",
		},
		tempUnit: Kelvin,
		convTempFuncs: map[tempType]ConvFunc{
			Celsius:    func(in float64) float64 { return in - 273.15 },
			Fahrenheit: func(in float64) float64 { return (in-273.15)*1.8 + 32 },
		},
	},
}
