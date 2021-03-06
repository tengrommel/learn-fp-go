// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Complex128

package num

// Complex128Slice is a slice of type Complex128. Use it where you would use []Complex128.
type Complex128Slice []Complex128

// SumComplex128 sums Complex128 over elements in Complex128Slice. See: http://clipperhouse.github.io/gen/#Sum
func (rcv Complex128Slice) SumComplex128(fn func(Complex128) Complex128) (result Complex128) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}
