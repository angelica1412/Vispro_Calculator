// calculator.go (sebagai implementasi dari file Operations)

package calculator

// Mengimport function yang ada pada mathops ke dalam package calculator
import (
	"Calculator_GoLang/mathops"
	"errors"
	"math"
)

type Calculator struct{}

// Function yang digunakan apabila terdapat 2 angka yang ingin dijumlahkan
func (c *Calculator) Add(a, b float64) float64 {
	return mathops.Addition(a, b)
}

// Function yang digunakan apabila terdapat 2 angka yang ingin dikurangkan
func (c *Calculator) Subtract(a, b float64) float64 {
	return mathops.Subtraction(a, b)
}

// Function yang digunakan apabila terdapat 2 angka yang ingin dikalikan
func (c *Calculator) Multiply(a, b float64) float64 {
	return mathops.Multiplication(a, b)
}

// Function yang digunakan apabila terdapat 2 angka yang ingin dibagi
func (c *Calculator) Divide(a, b float64) (float64, error) {
	result, err := mathops.Division(a, b)

	// jika penyebutnya adalah 0, akan menampilkan eror
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Function yang digunakan untuk menghitung hasil pemangkatan suatu bilangan dengan eksponen tertentu.
func (c *Calculator) Power(base, exponent float64) (float64, error) {
	return mathops.Power(base, exponent)
}

// Function untuk menghitung akar kuadrat dari suatu angka.
func (c *Calculator) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative number is undefined")
	}
	return mathops.SquareRoot(a)
}

// Function untuk menghitung sinus dari suatu sudut (dalam derajat).
func (c *Calculator) Sin(angle float64) float64 {
	// Konversi derajat ke radian
	angleInRadians := angle * math.Pi / 180.0
	return mathops.Sin(angleInRadians)
}

// Function untuk menghitung cosinus dari suatu sudut (dalam derajat).
func (c *Calculator) Cos(angle float64) float64 {
	// Konversi derajat ke radian
	angleInRadians := angle * math.Pi / 180.0
	return mathops.Cos(angleInRadians)
}

// Function untuk menghitung tangen dari suatu sudut (dalam derajat).
func (c *Calculator) Tan(angle float64) (float64, error) {
	// Convert degrees to radians
	angleInRadians := angle * math.Pi / 180.0

	cosValue := mathops.Cos(angleInRadians)

	// eror yang akan dijalankan jika hasil dari cos sudut yang diinput oleh user bernilai 0
	if math.Abs(cosValue) < 1e-10 { // 1e-10 merupakan nilai epsilon yang merupakan bilangan yang relatif kecil
		return math.Inf(1), errors.New("tangent is undefined for cos(θ) ≈ 0")
	}

	return mathops.Sin(angleInRadians) / cosValue, nil
}
