// ============================================================================
// = builder.go
// = 	Description		Builds/converts things for big.Int.
// = 	Date			December 09, 2021
// ============================================================================

package utils

// converts from []int64 slice to []*big.Int
func ToBigSlice(slice []int64) []*bint {
	a := iSlice(int64(len(slice)))
	for i := 0; i < len(slice); i++ {
		a[i] = inew(slice[i])
	}
	return a
}

// converts from []*big.Int to []int64
func ToIntSlice(slice []*bint) []int64 {
	a := make([]int64, 0)
	for i := 0; i < len(slice); i++ {
		a = append(a, slice[i].Int64())
	}
	return a
}
