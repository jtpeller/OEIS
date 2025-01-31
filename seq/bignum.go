// ============================================================================
// = bignum.go
// = 	Description		math/big wrappers for readability
// = 	Date			December 09, 2021
// ============================================================================

package seq

import (
	"math/big"

	gb "github.com/jtpeller/gobig"
)

// BIG.INT CALCULATIONS
func zero() *big.Int {
	return gb.Zero()
}

func inew(i int64) *big.Int {
	return gb.New(i)
}

func abs(a *big.Int) *big.Int {
	return gb.Abs(a)
}

func add(a, b *big.Int) *big.Int {
	return gb.Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return gb.Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return gb.Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return gb.Div(a, b)
}

func pow(a *big.Int, e *big.Int) *big.Int {
	return gb.Pow(a, e)
}

func sqrt(a *big.Int) *big.Int {
	return gb.Sqrt(a)
}

// ### comparisons
func equals(a, b *big.Int) bool {
	return gb.Equals(a, b)
}

func lt(a, b *big.Int) bool {
	return gb.Less(a, b)
}

func lteq(a, b *big.Int) bool {
	return gb.LessEqual(a, b)
}

func gt(a, b *big.Int) bool {
	return gb.Greater(a, b)
}

func gteq(a, b *big.Int) bool {
	return gb.GreaterEqual(a, b)
}

// BIG.FLOAT CALCULATIONS
func fzero() *big.Float {
	return gb.ZeroFloat()
}

func fnew(a float64) *big.Float {
	return gb.NewFloat(a)
}

func fabs(a *big.Float) *big.Float {
	return gb.AbsFloat(a)
}

func fadd(a, b *big.Float) *big.Float {
	return gb.AddFloat(a, b)
}

func fsub(a, b *big.Float) *big.Float {
	return gb.SubFloat(a, b)
}

func fmul(a, b *big.Float) *big.Float {
	return gb.MulFloat(a, b)
}

func fdiv(a, b *big.Float) *big.Float {
	return gb.DivFloat(a, b)
}

func fpow(a *big.Float, e int64) *big.Float {
	return gb.PowFloat(a, e)
}

func fsqrt(a *big.Float) *big.Float {
	return gb.SqrtFloat(a)
}

func floor(a *big.Float) *big.Int {
	return gb.Floor(a)
}

func tofloat(a *big.Int) *big.Float {
	return gb.ToFloat(a)
}

func round(a *big.Float) *big.Int {
	return gb.Round(a)
}
