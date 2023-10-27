package cake

func (s *Slice[T]) MapToString(cfn CallBackFn[T, string]) Slice[string] {
	return mapTo[string](s, cfn)
}
