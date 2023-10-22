package cake

import (
	"reflect"
	"testing"
)

type TestedType interface {
	int | string | any
}

type testCase[T TestedType] struct {
	name     string
	input    []T
	expected Slice[T]
}

func TestGenerics(t *testing.T) {
	testCases := []testCase[any]{
		{
			name:     "mix types case",
			input:    []any{1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}},
			expected: Slice[any]{1, "apple", 1.5, struct{}{}, []int{}, map[string]string{"1": "apple"}},
		},
		{
			name:     "empty slice case",
			input:    []any{},
			expected: Slice[any]{},
		},
		{
			name:     "large number slice case",
			input:    make([]any, 10000),
			expected: make(Slice[any], 10000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, runTestCase(tc))
	}
}

func TestInteger(t *testing.T) {
	testCases := []testCase[int]{
		{
			name:     "normal case",
			input:    []int{1, 2, 3},
			expected: Slice[int]{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, runTestCase(tc))
	}
}

func TestString(t *testing.T) {
	testCases := []testCase[string]{
		{
			name:     "normal case",
			input:    []string{"apple", "banana", "cherry"},
			expected: Slice[string]{"apple", "banana", "cherry"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, runTestCase(tc))
	}
}

func runTestCase[T TestedType](tc testCase[T]) func(t *testing.T) {
	return func(t *testing.T) {
		result := From(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("Expected length %d, got %d", len(tc.expected), len(result))
		}
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
