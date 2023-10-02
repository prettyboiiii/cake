package cake

/*
Making a new cake.Slice from value(s) of type T

Param:

	optional param(s): value(s) that will be in the new cake.Slice

Returns:

	the new cake.Slice[T] with the elements from param(s) in it
*/
func New[T any](elems ...T) Slice[T] {
	sc := Slice[T](elems)
	return sc
}
