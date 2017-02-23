package datatype

import "fmt"

type typeGroup int

//go:generate stringer -type=typeGroup
const (
	GroupBare typeGroup = iota
	GroupLength
	GroupWeight
	GroupTemperature
	GroupVolume
	GroupTime
	GroupCurrency
	GroupDataSize
)

type DataType struct {
	Group       typeGroup
	Names       []string
	DisplayName string
	Factor      float64
}

var BareDataType = &DataType{Group: GroupBare, Factor: 1}

func (typeFrom *DataType) GetConvFunc(typeTo *DataType) (func(float64) float64, error) {
	if typeFrom.Group == GroupBare {
		return func(in float64) float64 { return in }, nil
	} else if typeFrom.Group != typeTo.Group {
		return nil, fmt.Errorf("GetConversionMultipl: incompatible types %s - %s", typeFrom.DisplayName, typeTo.DisplayName)
	}

	return func(in float64) float64 { return in * typeTo.Factor / typeFrom.Factor }, nil
}

func GetType(name string) (*DataType, error) {

	if dt, ok := typeMap[name]; ok {
		return dt, nil
	}
	return &DataType{}, fmt.Errorf("getTyp: unknown datatype %q", name)
}

var typeMap = make(map[string]*DataType)

func initUnits(units []*DataType) {
	for _, t := range units {
		for _, n := range t.Names {
			typeMap[n] = t
		}
	}
}

func init() {
	initUnits(lengthTypes)
	initUnits(weightTypes)
	initUnits(volumeTypes)

}
