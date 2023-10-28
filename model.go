package cake

type Slice[T any] []T

type Predicate[T any] func(T) bool
