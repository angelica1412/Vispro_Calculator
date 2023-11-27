// calculator_test.go

package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	c := Calculator{}
	result := c.Add(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	c := Calculator{}
	result := c.Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %f", result)
	}
}

func TestMultiply(t *testing.T) {
	c := Calculator{}
	result := c.Multiply(3, 5)
	if result != 15 {
		t.Errorf("Expected 15, but got %f", result)
	}
}

func TestDivide(t *testing.T) {
	c := Calculator{}
	result, err := c.Divide(6, 2)
	if err != nil || result != 3 {
		t.Errorf("Expected result=3 and error=nil, but got result=%f and error=%v", result, err)
	}

	_, err = c.Divide(6, 0)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestPower(t *testing.T) {
	c := Calculator{}
	result, err := c.Power(2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSquareRoot(t *testing.T) {
	// Test case for a positive number
	calc := Calculator{}
	result, err := calc.SquareRoot(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("SquareRoot(4) = %v; want 2", result)
	}

	// Test case for a negative number
	result, err = calc.SquareRoot(-4)
	if err == nil {
		t.Error("Expected an error for SquareRoot(-4) but got none")
	}
}

// Toleransi kecil untuk membandingkan nilai floating-point
const epsilon = 1e-9

func TestSin(t *testing.T) {
	// Menguji Sin dengan nilai positif
	calc := Calculator{}
	result := calc.Sin(30)
	expected := 0.5
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Sin(30) = %v, expected %v", result, expected)
	}

	// Menguji Sin dengan nilai negatif
	result = calc.Sin(-45)
	expected = -0.7071067811865475
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Sin(-45) = %v, expected %v", result, expected)
	}

	// Menguji Sin dengan nilai nol
	result = calc.Sin(0)
	expected = 0
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Sin(0) = %v, expected %v", result, expected)
	}
}
func TestCos(t *testing.T) {
	// Menguji Cos dengan nilai positif
	calc := Calculator{}
	result := calc.Cos(60)
	expected := 0.5
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Cos(60) = %v, expected %v", result, expected)
	}

	// Menguji Cos dengan nilai negatif
	result = calc.Cos(-90)
	expected = 0
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Cos(-90) = %v, expected %v", result, expected)
	}

	// Menguji Cos dengan nilai nol
	result = calc.Cos(0)
	expected = 1
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Cos(0) = %v, expected %v", result, expected)
	}
}

func TestTan(t *testing.T) {
	// Menguji Tan dengan nilai positif
	calc := Calculator{}
	result, err := calc.Tan(60)
	expected := math.Sqrt(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Tan(60) = %v, expected %v", result, expected)
	}

	// Menguji Tan dengan nilai negatif
	result, err = calc.Tan(-45)
	expected = -1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Tan(-45) = %v, expected %v", result, expected)
	}

	// Menguji Tan dengan nilai nol
	result, err = calc.Tan(0)
	expected = 0
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if math.Abs(result-expected) > epsilon {
		t.Errorf("Tan(0) = %v, expected %v", result, expected)
	}
}
