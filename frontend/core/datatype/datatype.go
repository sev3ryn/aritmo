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

// ConvFunc - function that converts in value to out float. Specific for each datatype pair
type ConvFunc func(in float64) float64

// DataType - interface for maintaining datatype units manipulation - converting/displaying/mapping by name
type DataType interface {
	GetConvFunc(DataType) (ConvFunc, error)
	GetBase() *BaseDataType
}

// BaseDataType - basic unit info (names, displayName, unit group)
type BaseDataType struct {
	Group       typeGroup
	Names       []string
	DisplayName string
}

// String - stringer implementation
func (b *BaseDataType) String() string {
	return fmt.Sprintf("%s:%s", b.Group, b.DisplayName)
}

// SimpleDataType - DataType implementation that for units that could be converted by multyplying on coefficient
// it is related to such types as legth, area, volume, weight, etc.
type SimpleDataType struct {
	b      *BaseDataType
	Factor float64
}

// GetBase - return basic info(names, displayName, unit group) of datatype unit
func (t *SimpleDataType) GetBase() *BaseDataType {
	return t.b
}

// GetConvFunc - convert function for SimpleDataType. Finding unit-to-unit coefficient and multyplying on it
func (t *SimpleDataType) GetConvFunc(typeTo DataType) (ConvFunc, error) {
	if t.b.Group == GroupBare {
		return func(in float64) float64 { return in }, nil
	}

	typ, ok := typeTo.(*SimpleDataType)

	if !ok || t.b.Group != typ.b.Group {
		return nil, fmt.Errorf("GetConvFunc: incompatible types %#v - %#v", t, typeTo)
	}

	return func(in float64) float64 { return in * typ.Factor / t.Factor }, nil
}

// BareDataType - nil datatype - used for values where it is not specified
var BareDataType = &SimpleDataType{b: &BaseDataType{Group: GroupBare}, Factor: 1}

// GetType - type lookup by name
func (t TypeMap) GetType(name string) (DataType, error) {

	if dt, ok := t[name]; ok {
		return dt, nil
	}
	return nil, fmt.Errorf("GetType: unknown datatype %q", name)
}

// initUnits - datatype units initialization by adding them to datatype by name map
func (t TypeMap) initUnits(units []DataType) {
	for _, unit := range units {
		for _, name := range unit.GetBase().Names {
			t[name] = unit
		}
	}
}

// TypeMap - all supported datatype by name map
type TypeMap map[string]DataType

// Init - all avaliable units initialization - adding to TypeMap
func Init(currUpdateCh <-chan []byte) TypeMap {
	var typeMap = make(TypeMap)

	typeMap.initUnits(lengthTypes)
	typeMap.initUnits(weightTypes)
	typeMap.initUnits(volumeTypes)
	typeMap.initUnits(temperatureTypes)

	typeMap.initUnits(GetCurrUnits(currUpdateCh))

	return typeMap
}
