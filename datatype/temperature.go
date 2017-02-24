package datatype

import (
	"fmt"
)

type tempType int

const (
	Celsius tempType = iota
	Fahrenheit
	Kelvin
)

type TemperatureType struct {
	b             *BaseDataType
	tempUnit      tempType
	convTempFuncs map[tempType]ConvFunc
}

func (t *TemperatureType) GetBase() *BaseDataType {
	return t.b
}

func (typeFrom *TemperatureType) GetConvFunc(typeTo DataType) (ConvFunc, error) {
	if typeFrom.b.Group == GroupBare {
		return func(in float64) float64 { return in }, nil
	} else if typeFrom.b.Group != typeTo.GetBase().Group {
		return nil, fmt.Errorf("GetConversionMultipl: incompatible types %s - %s", typeFrom, typeTo)
	}

	return typeFrom.convTempFuncs[typeTo.(*TemperatureType).tempUnit], nil
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
