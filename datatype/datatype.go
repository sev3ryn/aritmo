package datatype

import "fmt"

type typeGroup int

//go:generate stringer -type=typeGroup
const (
	GroupBare typeGroup = iota
	GroupLength
	GroupWeight
	GroupTemperature
	GroupTime
	GroupVolume
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

func (typeFrom *DataType) GetConversionMultipl(typeTo *DataType) (float64, error) {
	if typeFrom.Group == GroupBare {
		return 1, nil
	} else if typeFrom.Group != typeTo.Group {
		return 0, fmt.Errorf("GetConversionMultipl: incompatible types %s - %s", typeFrom.DisplayName, typeTo.DisplayName)
	}

	return typeTo.Factor / typeFrom.Factor, nil
}

func GetType(name string) (*DataType, error) {

	if dt, ok := typeMap[name]; ok {
		return dt, nil
	}
	return &DataType{}, fmt.Errorf("getTyp: unknown datatype %q", name)
}

var typeMap = make(map[string]*DataType)

func init() {
	for _, t := range lengthTypes {
		for _, n := range t.Names {
			typeMap[n] = t
		}
	}
}
