package fungo

/*
	Software Engineer: Fachrin Aulia Nasution
	An always in beta software engineer
*/

import (
	"reflect"
)

type forEachCallback func(interface{}, int)
type filterCallback func(interface{}, int) bool
type mapCallback func(interface{}, int) interface{}
type reduceCallback func(int64, interface{}) int64

// Slice represents as generic type for every slice
type Slice struct {
	length int
	sType  reflect.Type
	values reflect.Value
}

// NewSlice will instantiate the new object of Slice
func NewSlice(slice interface{}) Slice {
	sliceValues := reflect.ValueOf(slice)
	sliceLength := sliceValues.Len()

	return Slice{
		length: sliceLength,
		values: sliceValues,
		sType:  sliceValues.Type(),
	}
}

// Map is the mapper
func (s Slice) Map(callback mapCallback) interface{} {
	if s.length == 0 {
		return nil
	}

	// get first element, create slice of it, and append it to slice
	var firstCallbackValue = reflect.ValueOf(callback(s.values.Index(0).Interface(), 0))
	var slice = createFormattedEmptyReflectionSlice(&s, firstCallbackValue.Interface())
	slice = appendSlice(&slice, &firstCallbackValue)

	if s.length > 1 {
		for index := 1; index < s.length; index++ {
			value := reflect.ValueOf(callback(s.values.Index(index).Interface(), index))
			slice = appendSlice(&slice, &value)
		}
	}

	return slice.Interface()
}

// Filter is the filter
func (s Slice) Filter(callback filterCallback) interface{} {
	var slice = createEmptyReflectionSlice(&s)

	for index := 0; index < s.length; index++ {
		value := s.values.Index(index)

		if callback(value.Interface(), index) {
			slice = appendSlice(&slice, &value)
		}
	}

	return slice.Interface()
}

// ForEach is the loop
func (s Slice) ForEach(callback forEachCallback) {
	for index := 0; index < s.length; index++ {
		callback(s.values.Index(index).Interface(), index)
	}
}

// Reduce is the reducer
func (s Slice) Reduce(callback reduceCallback) int64 {
	var accumulator = int64(0)

	for index := 0; index < s.length; index++ {
		value := s.values.Index(index)
		accumulator = callback(accumulator, value.Interface())
	}

	return accumulator
}

func createFormattedEmptyReflectionSlice(s *Slice, format interface{}) reflect.Value {
	var formatType = reflect.ValueOf(format).Type()
	var reflection = reflect.MakeSlice(reflect.SliceOf(formatType), 0, s.length)
	var reflectionValue = reflect.New(reflection.Type())
	return reflectionValue.Elem()
}

func createEmptyReflectionSlice(s *Slice) reflect.Value {
	var reflection = reflect.MakeSlice(reflect.SliceOf(s.sType.Elem()), 0, s.length)
	var reflectionValue = reflect.New(reflection.Type())
	return reflectionValue.Elem()
}

func appendSlice(slice *reflect.Value, value *reflect.Value) reflect.Value {
	slice.Set(reflect.Append(*slice, *value))
	return *slice
}
