package concurrent

import (
	"fmt"

	"github.com/skeptycal/gosimple/types/constraints"
)

type GetSetter[K constraints.Ordered, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V) error
}

type KVLister[K constraints.Ordered, V any] interface {
	Keys() []K
	Values() []V
}

type Dict[K constraints.Ordered, V any] interface {
	GetSetter[K, V]
	KVLister[K, V]
}

type dict[K constraints.Ordered, V any] struct {
	locked bool
	m      map[K]V
}

func (d *dict[K, V]) Get(key K) (value V, ok bool) {
	if v, ok := d.m[key]; ok {
		return v, true
	}
	return
}

func (d *dict[K, V]) Set(key K, value V) error {
	if d.locked {
		if _, ok := d.m[key]; ok {
			return fmt.Errorf("cannot change value for existing key: %v", key)
		}
	}
	d.m[key] = value
	return nil
}

func (d *dict[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d *dict[K, V]) Values() []V {
	values := make([]V, 0, len(d.m))
	for _, v := range d.m {
		values = append(values, v)
	}
	return values
}
