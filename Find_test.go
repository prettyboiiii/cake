package cake_test

import (
	"reflect"
	"testing"

	"github.com/prettyboiiii/cake"
)

type TestedSliceFindType interface {
	int | string | any
}

type testSliceFindCase[T TestedSliceFindType] struct {
	name     string
	slice    cake.Slice[T]
	pc       cake.Predicate[T]
	expected T
}

func TestSliceFind(t *testing.T) {
	genericTestCases := []testSliceFindCase[any]{
		{
			name:     "found",
			slice:    cake.New[any](1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}),
			expected: "apple",
			pc: func(v any) bool {
				return v == "apple"
			},
		},
		{
			name:     "not found",
			slice:    cake.New[any](1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}),
			expected: nil,
			pc: func(v any) bool {
				return v == 1.999
			},
		},
	}

	intergerTestCases := []testSliceFindCase[int]{
		{
			name:     "integer found",
			slice:    cake.New[int](1, 2, 3),
			expected: 3,
			pc: func(v int) bool {
				return v == 3
			},
		},
		{
			name:     "integer not found",
			slice:    cake.New[int](1, 2, 3),
			expected: 0,
			pc: func(v int) bool {
				return v == 99
			},
		},
	}

	stringTestCases := []testSliceFindCase[string]{
		{
			name:     "string found",
			slice:    cake.New[string]("apple", "banana", "cherry"),
			expected: "cherry",
			pc: func(v string) bool {
				return v == "cherry"
			},
		},
		{
			name:     "string not found",
			slice:    cake.New[string]("apple", "banana", "cherry"),
			expected: "",
			pc: func(v string) bool {
				return v == "UNKNOWN"
			},
		},
	}

	for _, tc := range genericTestCases {
		t.Run(tc.name, runTestSliceFind(tc))
	}

	for _, tc := range intergerTestCases {
		t.Run(tc.name, runTestSliceFind(tc))
	}

	for _, tc := range stringTestCases {
		t.Run(tc.name, runTestSliceFind(tc))
	}

}

func runTestSliceFind[T TestedSliceFindType](tc testSliceFindCase[T]) func(t *testing.T) {
	return func(t *testing.T) {
		result := tc.slice.Find(tc.pc)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
