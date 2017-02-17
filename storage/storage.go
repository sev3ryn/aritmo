package storage

import "fmt"

type StoreItem struct {
	val float64
	ok  bool
}

type Store map[string]StoreItem

var RAMStore = Store{}

func WrapStoreItem(v float64, err error) StoreItem {
	if err != nil {

		return StoreItem{ok: false}
	}
	return StoreItem{val: v, ok: true}
}

func (s Store) Get(k string) (float64, error) {
	v, ok := s[k]
	if !ok {
		return 0, fmt.Errorf("No such variable ", k)
	} else if !v.ok {
		return 0, fmt.Errorf("Invalid variable ", k)
	}
	return v.val, nil
}
