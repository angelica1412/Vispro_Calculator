package mathops

import (
	"errors"
	"math"
)

// Function untuk mengembalikan hasil penjumlahan
func Addition(a, b float64) float64 {
	return a + b
}

// Function untuk mengembalikan hasil pengurangan
func Subtraction(a, b float64) float64 {
	return a - b
}

// Function untuk mengembalikan hasil perkalian
func Multiplication(a, b float64) float64 {
	return a * b
}

// Function untuk mengembalikan hasil pembagian
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is undefined") // pesan eror yang akan muncul jika b (penyebutnya) adalah 0
	}
	return a / b, nil
}

// Function untuk mengembalikan hasil pemangkatan
func Power(base, exponent float64) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("undefined for 0 base with negative exponent") // pesan eror yang keluar jika basisnya adalah 0 dan pangkatnya berupa minus
	}
	return math.Pow(base, exponent), nil
}

// Function untuk mengembalikan hasil akar kuadrat
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative number is undefined") // pesan eror yang akan muncul ketika angka di dalam akar berupa nilai negatif
	}
	return math.Sqrt(a), nil
}

// Function untuk mengembalikan nilai sin dari sudut yang diberikan dalam bentuk radian
func Sin(angle float64) float64 {
	return math.Sin(angle)
}

// Function untuk mengembalikan nilai cos dari sudut yang diberikan dalam bentuk radian
func Cos(angle float64) float64 {
	return math.Cos(angle)
}

// Function untuk mengembalikan nilai tan dari sudut yang diberikan dalam bentuk radian
func Tan(angle float64) float64 {
	return math.Tan(angle)
}
