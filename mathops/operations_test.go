// operations_test.go

package mathops

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(3, 5)
	if result != 15 {
		t.Errorf("Expected 15, but got %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(6, 2)
	if err != nil || result != 3 {
		t.Errorf("Expected result=3 and error=nil, but got result=%f and error=%v", result, err)
	}

	_, err = Division(6, 0)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
