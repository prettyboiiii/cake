package sliceOfCake

import (
	"reflect"
)

type Predicate[T any] func(T) bool

type sliceOfCake[T any] []T

type SliceOfCake interface {
	find(arr []any, predicate Predicate[any], result interface{})
}

func New[T any, F sliceOfCake[T]](elems ...T) *sliceOfCake[T] {
	sc := sliceOfCake[T](elems)
	return &sc
}

func (c sliceOfCake[T]) Find(predicate Predicate[T]) *T {
	arrValue := reflect.ValueOf(c)

	if arrValue.Kind() != reflect.Slice {
		panic("arr must be a slice")
	}

	for _, v := range c {
		if predicate(v) {
			return &v
		}
	}

	return nil
}

func (c *sliceOfCake[T]) Append(elems ...T) int {
	*c = append(*c, elems...)
	return len(*c)
}
