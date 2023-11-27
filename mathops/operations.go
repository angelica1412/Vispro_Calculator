// operations.go

package mathops

import (
	"errors"
	"math"
)

// Addition returns the sum of two numbers.
func Addition(a, b float64) float64 {
	return a + b
}

// Subtraction returns the difference between two numbers.
func Subtraction(a, b float64) float64 {
	return a - b
}

// Multiplication returns the product of two numbers.
func Multiplication(a, b float64) float64 {
	return a * b
}

// Division returns the result of dividing a by b.
// It returns an error if b is 0.
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is undefined")
	}
	return a / b, nil
}

// Power returns the result of base raised to the power of exponent.
// It returns an error if the base is 0 and the exponent is negative.
func Power(base, exponent float64) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("undefined for 0 base with negative exponent")
	}
	return math.Pow(base, exponent), nil
}

// SquareRoot returns the square root of a number.
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative number is undefined")
	}
	return math.Sqrt(a), nil
}

// Sin returns the sine of the given angle in radians.
func Sin(angle float64) float64 {
	return math.Sin(angle)
}

// Cos returns the cosine of the given angle in radians.
func Cos(angle float64) float64 {
	return math.Cos(angle)
}

// Tan returns the tangent of the given angle in radians.
func Tan(angle float64) float64 {
	return math.Tan(angle)
}
