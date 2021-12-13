// ============================================================================
// = bignum.go																  =
// = 	Description: uses my custom gobig package to handle big nums. I use	  =
// =			wrappers for readability	 								  =
// = 			e.g. big.NewInt(0) vs gb.New(0) vs New(0)					  =
// = 	Date: December 09, 2021												  =
// ============================================================================

package seq

import (
	"math/big"

	gb "github.com/jtpeller/gobig"
)

// BIG.INT CALCULATIONS
func Zero() *big.Int {
	return gb.Zero()
}

func New(i int64) *big.Int {
	return gb.New(i)
}

func Abs(a *big.Int) *big.Int {
	return gb.Abs(a)
}

func Add(a, b *big.Int) *big.Int {
	return gb.Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return gb.Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return gb.Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return gb.Div(a, b)
}

func Pow(a *big.Int, e *big.Int) *big.Int {
	return gb.Pow(a, e)
}

func Sqrt(a *big.Int) *big.Int {
	return gb.Sqrt(a)
}

func Lsh(a *big.Int, e int) *big.Int {
	return gb.Lsh(a, e)
}

// BIG.FLOAT CALCULATIONS
func ZeroFloat() *big.Float {
	return gb.ZeroFloat()
}

func NewFloat(a float64) *big.Float {
	return gb.NewFloat(a)
}

func AbsFloat(a *big.Float) *big.Float {
	return gb.AbsFloat(a)
}

func AddFloat(a, b *big.Float) *big.Float {
	return gb.AddFloat(a, b)
}

func SubFloat(a, b *big.Float) *big.Float {
	return gb.SubFloat(a, b)
}

func MulFloat(a, b *big.Float) *big.Float {
	return gb.MulFloat(a, b)
}

func DivFloat(a, b *big.Float) *big.Float {
	return gb.DivFloat(a, b)
}

func PowFloat(a *big.Float, e int64) *big.Float {
	return gb.PowFloat(a, e)
}

func SqrtFloat(a *big.Float) *big.Float {
	return gb.SqrtFloat(a)
}

func Floor(a *big.Float) *big.Int {
	return gb.Floor(a)
}

func ToFloat(a *big.Int) *big.Float {
	return gb.ToFloat(a)
}

func Round(a *big.Float) *big.Int {
	return gb.Round(a)
}