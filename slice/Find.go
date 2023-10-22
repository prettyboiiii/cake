package cake

/*
Find the particular value follow the predicate func in the cake.Slice
Param:

  - param1: predicate func, its return value indicate what value to find

Returns:

	The founded value (`pc` returns `true`)
	if the value does not exist in the cake.Slice, zero value of T is sent
*/
func (s Slice[T]) Find(pc Predicate[T]) T {
	for _, v := range s {
		if pc(v) {
			return v
		}
	}
	var zeroV T

	return zeroV
}
