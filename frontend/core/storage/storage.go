package storage

import (
	"fmt"

	"github.com/sev3ryn/aritmo/frontend/core/datatype"
)

// Result - the smallest unit of parse operations. Consists of value and datatype
type Result struct {
	Val float64
	Typ datatype.DataType
}

// Store - interface for variable storage
type Store interface {
	Save(key string, val Result, err error) error
	Get(key string) (val Result, err error)
	Remove(key string) error
}

// StoreItem - atomic type of storage
type StoreItem struct {
	line int
	val  Result
	ok   bool
}

// StoreItemMap - implementation of Store in RAM
type StoreItemMap struct {
	m        map[string][]StoreItem
	CurrLine int
}

// RAMStore - instance of StoreItemMap
var RAMStore = StoreItemMap{m: make(map[string][]StoreItem)}

// Insert - add variable to list of variable declarations according to its postion in editor
func Insert(slice []StoreItem, index int, value StoreItem) []StoreItem {
	// Grow the slice by one element.
	slice = append(slice, slice[0])
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}

func (s StoreItemMap) save(key string, item StoreItem) {
	vars, ok := s.m[key]
	if !ok {
		s.m[key] = []StoreItem{item}
		return
	}
	for i := 0; i < len(vars); i++ {
		if vars[i].line == item.line {
			// update item
			s.m[key][i] = item
			return
		} else if vars[i].line > item.line {
			// insert item into current position
			s.m[key] = Insert(s.m[key], i, item)
			return
		}
	}

	s.m[key] = append(s.m[key], item)

}

// Save - store variable
func (s StoreItemMap) Save(key string, val Result, err error) error {
	if err != nil {
		s.save(key, StoreItem{line: s.CurrLine, ok: false})
		return nil
	}
	s.save(key, StoreItem{line: s.CurrLine, val: val, ok: true})
	return nil
}

// Get - get nearest(if multiple declaration) variable by key
func (s StoreItemMap) Get(k string) (Result, error) {
	v, ok := s.m[k]

	fmt.Println(s.CurrLine)
	if !ok || v[0].line > s.CurrLine {
		return Result{}, fmt.Errorf("No such variable %s", k)
	}
	//reverse search as for most cases it will require less operations
	for i := len(v) - 1; i >= 0; i-- {
		if v[i].line < s.CurrLine {
			if !v[i].ok {
				return Result{}, fmt.Errorf("Invalid variable %s", k)
			}
			return v[i].val, nil
		}
	}
	return Result{}, fmt.Errorf("storage.Get: Impossible case")
}

// Remove - remove key from storage. Not used now
func (s StoreItemMap) Remove(k string) error {
	//s[k] = StoreItem{}
	return nil
}
