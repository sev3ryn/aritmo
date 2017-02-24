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

type ConvFunc func(in float64) float64

type DataType interface {
	GetConvFunc(DataType) (ConvFunc, error)
	GetBase() *BaseDataType
}

type BaseDataType struct {
	Group       typeGroup
	Names       []string
	DisplayName string
}

func (b *BaseDataType) String() string {
	return fmt.Sprintf("%s:%s", b.Group, b.DisplayName)
}

func (b *BaseDataType) GetNames() []string {
	return b.Names
}

type SimpleDataType struct {
	b      *BaseDataType
	Factor float64
}

func (t *SimpleDataType) GetBase() *BaseDataType {
	return t.b
}

func (typeFrom *SimpleDataType) GetConvFunc(typeTo DataType) (ConvFunc, error) {
	if typeFrom.b.Group == GroupBare {
		return func(in float64) float64 { return in }, nil
	}

	typ, ok := typeTo.(*SimpleDataType)

	if !ok || typeFrom.b.Group != typ.b.Group {
		return nil, fmt.Errorf("GetConvFunc: incompatible types %s - %s", typeFrom, typeTo)
	}

	return func(in float64) float64 { return in * typ.Factor / typeFrom.Factor }, nil
}

var BareDataType = &SimpleDataType{b: &BaseDataType{Group: GroupBare}, Factor: 1}

func (t TypeMap) GetType(name string) (DataType, error) {

	if dt, ok := t[name]; ok {
		return dt, nil
	}
	return nil, fmt.Errorf("getTyp: unknown datatype %q", name)
}

func (t TypeMap) initUnits(units []DataType) {
	for _, unit := range units {
		for _, name := range unit.GetBase().GetNames() {
			t[name] = unit
		}
	}
}

type TypeMap map[string]DataType

func Init() TypeMap {
	var typeMap = make(TypeMap)

	typeMap.initUnits(lengthTypes)
	typeMap.initUnits(weightTypes)
	typeMap.initUnits(volumeTypes)
	typeMap.initUnits(temperatureTypes)
	return typeMap
}
