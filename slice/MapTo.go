package cake

func mapTo[R any, T any](s *Slice[T], cfn CallBackFn[T, R]) Slice[R] {
	mapped := New[R]()
	for i, v := range *s {
		mapped.Push(cfn(i, v))
	}
	return mapped
}

func MapTo[R any, T any](s *Slice[T], cfn CallBackFn[T, R]) Slice[R] {
	return mapTo(s, cfn)
}
