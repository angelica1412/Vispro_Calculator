// main.go

package main

import (
	"Calculator_GoLang/calculator"
	"Calculator_GoLang/ui"
)

func main() {
	// Inisialisasi kalkulator dan antarmuka pengguna
	calc := calculator.Calculator{}  // Membuat instance dari kalkulator
	ui := ui.NewUserInterface(&calc) // Membuat instance dari antarmuka pengguna dengan mengirimkan referensi kalkulator

	// Menjalankan kalkulator
	ui.RunCalculator() // Memulai antarmuka pengguna untuk menjalankan kalkulator
}
