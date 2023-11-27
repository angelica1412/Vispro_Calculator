// calculator.go (sebagai implementasi dari file Operations)

package calculator

import (
	"Calculator_GoLang/mathops"
	"errors"
	"math"
)

// Calculator struct represents a simple calculator.
type Calculator struct{}

// Add performs addition of two numbers.
func (c *Calculator) Add(a, b float64) float64 {
	return mathops.Addition(a, b)
}

// Subtract performs subtraction between two numbers.
func (c *Calculator) Subtract(a, b float64) float64 {
	return mathops.Subtraction(a, b)
}

// Multiply performs multiplication between two numbers.
func (c *Calculator) Multiply(a, b float64) float64 {
	return mathops.Multiplication(a, b)
}

// Divide performs division between two numbers.
func (c *Calculator) Divide(a, b float64) (float64, error) {
	result, err := mathops.Division(a, b)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Power calculates the power of a number to the specified exponent.
func (c *Calculator) Power(base, exponent float64) (float64, error) {
	return mathops.Power(base, exponent)
}

// SquareRoot calculates the square root of a number.
func (c *Calculator) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative number is undefined")
	}
	return mathops.SquareRoot(a)
}

func (c *Calculator) Sin(angle float64) float64 {
	// Konversi derajat ke radian
	angleInRadians := angle * math.Pi / 180.0
	return math.Sin(angleInRadians)
}

// Cos calculates the cosine of an angle (in degrees).
func (c *Calculator) Cos(angle float64) float64 {
	// Konversi derajat ke radian
	angleInRadians := angle * math.Pi / 180.0
	return math.Cos(angleInRadians)
}

// Tan calculates the tangent of an angle (in degrees).
func (c *Calculator) Tan(angle float64) (float64, error) {
	// Konversi derajat ke radian
	angleInRadians := angle * math.Pi / 180.0

	// Tangen dari sudut theta adalah sin(theta) / cos(theta)
	cosValue := math.Cos(angleInRadians)
	if cosValue == 0 {
		return 0, errors.New("tangent is undefined for cos(θ) = 0")
	}

	return math.Sin(angleInRadians) / cosValue, nil
}
