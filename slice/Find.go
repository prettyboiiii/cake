package cake

/*
Find is a sample function that takes two parameters and returns a result.
Param:

	- param1: predicate func, its return value indicate what value to find

Returns:

	founded value (`pc` returns `true`) if it existed in the Slice, otherwise zero value of T is returned
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
