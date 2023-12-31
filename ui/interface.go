// interface.go

package ui

import (
	"Calculator_GoLang/calculator"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// UserInterface adalah struktur yang mewakili antarmuka pengguna untuk kalkulator.
type UserInterface struct {
	calculator *calculator.Calculator
}

// Function untuk membuat instance user interface baru
func NewUserInterface(calc *calculator.Calculator) *UserInterface {
	return &UserInterface{calculator: calc}
}

// Function yang dijalankan untuk meminta pengguna untuk input dan mengembalikan nilai yang dimasukkan.
func (ui *UserInterface) getUserInput(message string) (float64, error) {
	fmt.Print(message + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	// Jika inoutan user berupa spasiS
	if input == "" {
		return 0, errors.New("input is empty")
	}

	// Validasi apakah input hanya mengandung angka atau angka desimal
	dotCount := 0
	for i, char := range input {
		if i == 0 && char == '-' {
			continue
		}

		// Jika inputan user berupa huruf
		if char == '.' {
			// Hanya izinkan satu titik desimal
			if dotCount > 0 {
				return 0, errors.New("invalid input: must be a number")
			}
			dotCount++
			continue
		}
		if !unicode.IsDigit(char) {
			return 0, errors.New("invalid input: must be a number")
		}
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %v", err)
	}

	return value, nil
}

func (ui *UserInterface) displayResult(result float64, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		// Gunakan %g atau %f untuk menyesuaikan jumlah desimal secara otomatis.
		fmt.Printf("Result: %.1f\n", result)
	}
}

// Function untuk meminta pengguna untuk input dan melakukan operasi kalkulator.
func (ui *UserInterface) RunCalculator() {
	for {
		fmt.Println(" ")
		fmt.Print("=================")
		fmt.Println("\nCalculator Menu:")
		fmt.Println("=================")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Power")
		fmt.Println("6. Square Roots")
		fmt.Println("7. Trigonometri")
		fmt.Println("8. Exit")

		choice, err := ui.getUserInput("Enter your choice")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Println(" ")

		switch choice {

		// Akan jalan sesuai dengan inputan user
		case 1:
			a, err := ui.getUserInput("Enter the first number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			b, err := ui.getUserInput("Enter the second number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			result := ui.calculator.Add(a, b)
			ui.displayResult(result, nil)
		case 2:
			a, err := ui.getUserInput("Enter the first number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			b, err := ui.getUserInput("Enter the second number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			result := ui.calculator.Subtract(a, b)
			ui.displayResult(result, nil)
		case 3:
			a, err := ui.getUserInput("Enter the first number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			b, err := ui.getUserInput("Enter the second number")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			result := ui.calculator.Multiply(a, b)
			ui.displayResult(result, nil)
		case 4:
			a, err := ui.getUserInput("Enter the dividend")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			b, err := ui.getUserInput("Enter the divisor")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			result, err := ui.calculator.Divide(a, b)
			ui.displayResult(result, err)
		case 5:
			base, err := ui.getUserInput("Enter the base")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			exponent, err := ui.getUserInput("Enter the exponent")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			result, err := ui.calculator.Power(base, exponent)
			ui.displayResult(result, err)

		case 6:
			a, _ := ui.getUserInput("Enter the number to find the square root")
			result, err := ui.calculator.SquareRoot(a)
			ui.displayResult(result, err)

		case 7:
			angle, err := ui.getUserInput("Enter the angle (in degrees) for trigonometric functions")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			// Pilih fungsi trigonometri
			fmt.Println("Trigonometric Functions:")
			fmt.Println("1. Sin")
			fmt.Println("2. Cos")
			fmt.Println("3. Tan")

			trigChoice, err := ui.getUserInput("Enter your choice for trigonometric function")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			// Bagian yang akan dijalankan berdasarkan inputan (pilihan) dari user
			switch trigChoice {
			case 1:
				result := ui.calculator.Sin(angle)
				ui.displayResult(result, nil)
			case 2:
				result := ui.calculator.Cos(angle)
				ui.displayResult(result, nil)
			case 3:
				result, err := ui.calculator.Tan(angle)
				ui.displayResult(result, err)
			default:
				fmt.Println("Invalid choice for trigonometric function. Please enter a valid option.")
			}

		case 8:
			fmt.Println("Exiting Calculator. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")

		}
	}
}
