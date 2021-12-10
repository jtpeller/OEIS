// ============================================================================
// = bignum.go																  =
// = 	Description: VERY helpful wrapper functions for big.Int and big.Float =
// = 	Date: December 09, 2021												  =
// ============================================================================

package seq

import "math/big"

// BIG.INT CALCULATIONS
func Zero() *big.Int {
	r := big.NewInt(0)
	return r
}

func New(i int64) *big.Int {
	r := big.NewInt(i)
	return r
}

func Add(a, b *big.Int) *big.Int {
	return Zero().Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return Zero().Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return Zero().Mul(a,b)
}

func Div(a, b *big.Int) *big.Int {
	return Zero().Div(a, b)
}

func Pow(a *big.Int, e *big.Int) *big.Int {
	return Zero().Exp(a, e, big.NewInt(0))
}

// BIG.FLOAT CALCULATIONS
func ZeroFloat() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func NewFloat(a float64) *big.Float {
	r := big.NewFloat(a)
	r.SetPrec(256)
	return r
}

func AddFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Add(a, b)
}

func SubFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Sub(a, b)
}

func MulFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Mul(a,b)
}

func DivFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Quo(a, b)
}

func PowFloat(a *big.Float, e int64) *big.Float {
	r := ZeroFloat().Copy(a)
	for i := int64(0); i < e-1; i++ {
		r = MulFloat(r, a)
	}
	return r
}

func Truncate(a *big.Float) *big.Int {
	f, _ := a.Int(nil)
	return f
}

func BigIntToBigFloat(a *big.Int) *big.Float {
	r := new(big.Float)
	r.SetInt(a)
	r.SetPrec(256)
	return r
}

func RoundFloat(a *big.Float) *big.Int {
	a.Add(a, NewFloat(0.5))
	r, _ := a.Int(nil)
	return r
}