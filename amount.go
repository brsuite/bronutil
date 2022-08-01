// Copyright (c) 2013, 2014 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bronutil

import (
	"errors"
	"math"
	"strconv"
)

// AmountUnit describes a method of converting an Amount to something
// other than the base unit of a brocoin.  The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in brocoin to an amount counted in units.
type AmountUnit int

// These constants define various units used when describing a brocoin
// monetary amount.
const (
	AmountMegaBRON  AmountUnit = 6
	AmountKiloBRON  AmountUnit = 3
	AmountBRON      AmountUnit = 0
	AmountMilliBRON AmountUnit = -3
	AmountMicroBRON AmountUnit = -6
	AmountBronees   AmountUnit = -8
)

// String returns the unit as a string.  For recognized units, the SI
// prefix is used, or "Bronees" for the base unit.  For all unrecognized
// units, "1eN BRON" is returned, where N is the AmountUnit.
func (u AmountUnit) String() string {
	switch u {
	case AmountMegaBRON:
		return "MBRON"
	case AmountKiloBRON:
		return "kBRON"
	case AmountBRON:
		return "BRON"
	case AmountMilliBRON:
		return "mBRON"
	case AmountMicroBRON:
		return "μBRON"
	case AmountBronees:
		return "Bronees"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " BRON"
	}
}

// Amount represents the base brocoin monetary unit (colloquially referred
// to as a `Bronees').  A single Amount is equal to 1e-8 of a brocoin.
type Amount int64

// round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating point value representing
// some value in brocoin.  NewAmount errors if f is NaN or +-Infinity, but
// does not check that the amount is within the total amount of brocoin
// producible as f may not refer to an amount at a single moment in time.
//
// NewAmount is for specifically for converting BRON to Bronees.
// For creating a new Amount with an int64 value which denotes a quantity of Bronees,
// do a simple type conversion from type int64 to Amount.
// See GoDoc for example: http://godoc.org/github.com/brsuite/bronutil#example-Amount
func NewAmount(f float64) (Amount, error) {
	// The amount is only considered invalid if it cannot be represented
	// as an integer type.  This may happen if f is NaN or +-Infinity.
	switch {
	case math.IsNaN(f):
		fallthrough
	case math.IsInf(f, 1):
		fallthrough
	case math.IsInf(f, -1):
		return 0, errors.New("invalid brocoin amount")
	}

	return round(f * BroneesperBrocoin), nil
}

// ToUnit converts a monetary amount counted in brocoin base units to a
// floating point value representing an amount of brocoin.
func (a Amount) ToUnit(u AmountUnit) float64 {
	return float64(a) / math.Pow10(int(u+8))
}

// ToBRON is the equivalent of calling ToUnit with AmountBRON.
func (a Amount) ToBRON() float64 {
	return a.ToUnit(AmountBRON)
}

// Format formats a monetary amount counted in brocoin base units as a
// string for a given unit.  The conversion will succeed for any unit,
// however, known units will be formated with an appended label describing
// the units with SI notation, or "Bronees" for the base unit.
func (a Amount) Format(u AmountUnit) string {
	units := " " + u.String()
	return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + units
}

// String is the equivalent of calling Format with AmountBRON.
func (a Amount) String() string {
	return a.Format(AmountBRON)
}

// MulF64 multiplies an Amount by a floating point value.  While this is not
// an operation that must typically be done by a full node or wallet, it is
// useful for services that build on top of brocoin (for example, calculating
// a fee by multiplying by a percentage).
func (a Amount) MulF64(f float64) Amount {
	return round(float64(a) * f)
}
