package cake

type CallBackFn[T any, R any] func(int, T) R

func (s *Slice[T]) MapToAny(cfn CallBackFn[T, any]) Slice[any] {
	return mapTo(s, cfn)
}
