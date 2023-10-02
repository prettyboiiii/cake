package cake

/*
Copy an existing traditional Go slice into a new cake.Slice.

Param:

	- param1: a slice of type T

Returns:

	The cake.Slice[T] with the value as the original slice
*/
func From[T any](sl []T) Slice[T] {
	return sl
}
