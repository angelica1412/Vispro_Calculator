// calculator_test.go

package calculator

import (
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
