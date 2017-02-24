package storage

import (
	"fmt"

	"github.com/sev3ryn/aritmo/datatype"
)

type Result struct {
	Val float64
	Typ datatype.DataType
}

type Store interface {
	Save(key string, val Result, err error) error
	Get(key string) (val Result, err error)
	Remove(key string) error
}

type StoreItem struct {
	line int
	val  Result
	ok   bool
}

type StoreItemMap struct {
	m        map[string][]StoreItem
	CurrLine int
}

var RAMStore = StoreItemMap{m: make(map[string][]StoreItem)}

func Insert(slice []StoreItem, index int, value StoreItem) []StoreItem {
	// Grow the slice by one element.
	fmt.Println("e")
	slice = append(slice, slice[0])
	fmt.Println("expanded")
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	fmt.Println("copied")
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

func (s StoreItemMap) Save(key string, val Result, err error) error {
	if err != nil {
		s.save(key, StoreItem{line: s.CurrLine, ok: false})
		return nil
	}
	fmt.Printf("Prev %+v\n", s.m)
	s.save(key, StoreItem{line: s.CurrLine, val: val, ok: true})
	fmt.Printf("Next %+v\n", s.m)
	return nil
}

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

func (s StoreItemMap) Remove(k string) error {
	//s[k] = StoreItem{}
	return nil
}
