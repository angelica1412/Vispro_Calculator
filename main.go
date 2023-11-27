// main.go

package main

import (
	"Calculator_GoLang/calculator"
	"Calculator_GoLang/ui"
)

func main() {
	// Inisialisasi kalkulator dan antarmuka pengguna
	calc := calculator.Calculator{}
	ui := ui.NewUserInterface(&calc)

	// Menjalankan kalkulator
	ui.RunCalculator()
}
