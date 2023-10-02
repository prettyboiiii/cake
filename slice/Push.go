package cake

/*
Adding new value(s) at the end of the cake.Slice just like append() func

Param:

	- param(s): value(s) that will be added into the cake.Slice

*/
func (s *Slice[T]) Push(elems ...T) {
	*s = append(*s, elems...)
}
