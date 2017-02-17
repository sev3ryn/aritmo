package storage

import "fmt"

type Store interface {
	Save(key string, val float64, err error) error
	Get(key string) (val float64, err error)
	Remove(key string) error
}

type StoreItem struct {
	val float64
	ok  bool
}

type StoreItemMap map[string]StoreItem

var RAMStore = StoreItemMap{}

func (s StoreItemMap) Save(key string, val float64, err error) error {
	if err != nil {
		s[key] = StoreItem{ok: false}
		return nil
	}
	s[key] = StoreItem{val: val, ok: true}
	return nil
}

func (s StoreItemMap) Get(k string) (float64, error) {
	v, ok := s[k]
	if !ok {
		return 0, fmt.Errorf("No such variable %s", k)
	} else if !v.ok {
		return 0, fmt.Errorf("Invalid variable %s", k)
	}
	return v.val, nil
}

func (s StoreItemMap) Remove(k string) error {
	s[k] = StoreItem{}
	return nil
}
