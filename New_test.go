package cake_test

import (
	"reflect"
	"testing"

	"github.com/prettyboiiii/cake"
)

type TestedSliceNewType interface {
	int | string | any
}

type testSliceNewCase[T TestedSliceNewType] struct {
	name     string
	input    []T
	expected cake.Slice[T]
}

func TestSliceNew(t *testing.T) {
	genericTestCases := []testSliceNewCase[any]{
		{
			name:     "mix types case",
			input:    []any{1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}},
			expected: cake.Slice[any]{1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}},
		},
		{
			name:     "empty slice case",
			input:    []any{},
			expected: cake.Slice[any]{},
		},
		{
			name:     "large number slice case",
			input:    make([]any, 10000),
			expected: make(cake.Slice[any], 10000),
		},
	}

	integerTestCases := []testSliceNewCase[int]{
		{
			name:     "integer case",
			input:    []int{1, 2, 3},
			expected: cake.Slice[int]{1, 2, 3},
		},
	}

	stringTestCases := []testSliceNewCase[string]{
		{
			name:     "string case",
			input:    []string{"apple", "banana", "cherry"},
			expected: cake.Slice[string]{"apple", "banana", "cherry"},
		},
	}

	for _, tc := range genericTestCases {
		t.Run(tc.name, runTestSliceNew(tc))
	}

	for _, tc := range integerTestCases {
		t.Run(tc.name, runTestSliceNew(tc))
	}

	for _, tc := range stringTestCases {
		t.Run(tc.name, runTestSliceNew(tc))
	}
}

func runTestSliceNew[T TestedSliceNewType](tc testSliceNewCase[T]) func(t *testing.T) {
	return func(t *testing.T) {
		result := cake.New(tc.input...)
		if len(result) != len(tc.expected) {
			t.Errorf("Expected length %d, got %d", len(tc.expected), len(result))
		}
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
