package util

import (
	"fmt"
	"reflect"
)

// TODO 汎用的に使えるようにする[https://zenn.dev/jy8752/articles/757f37b3a7c7cd]
type ValueObject[T any] interface {
	Value() T
	Equals(other ValueObject[T]) bool
	String() string
}

type valueObject[T any] struct {
	value T
}

func NewValueObject[T any](v T) ValueObject[T] {
	return &valueObject[T]{value: v}
}

func (v *valueObject[T]) Value() T {
	return v.value
}

func (v *valueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(v.Value(), other.Value())
}

// いらんかも
func (v *valueObject[T]) String() string {
	return fmt.Sprintf("%v", v.value)
}
