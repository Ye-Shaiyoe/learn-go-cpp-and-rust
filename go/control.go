// control.go
// Penjelasan tentang struktur kontrol dalam Go: if, for, switch

package main

import (
	"fmt"
	"time"
)

func main() {
	// IF STATEMENT
	fmt.Println("=== IF STATEMENT ===")
	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D atau E")
	}

	// IF dengan statement pendek
	if nilai := 75; nilai >= 80 {
		fmt.Println("Lulus dengan nilai:", nilai)
	} else {
		fmt.Println("Tidak lulus dengan nilai:", nilai)
	}

	// FOR LOOP
	fmt.Println("\n=== FOR LOOP ===")

	// For klasik
	fmt.Println("For klasik:")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For seperti while
	fmt.Println("For seperti while:")
	count := 0
	for count < 3 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// For range
	fmt.Println("For range:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// SWITCH STATEMENT
	fmt.Println("\n=== SWITCH STATEMENT ===")
	hari := time.Now().Weekday()

	switch hari {
	case time.Saturday, time.Sunday:
		fmt.Println("Hari ini adalah akhir pekan")
	case time.Monday:
		fmt.Println("Hari ini adalah Senin, semangat!")
	default:
		fmt.Printf("Hari ini adalah hari kerja: %s\n", hari)
	}

	// Switch tanpa kondisi (seperti if-else bersarang)
	fmt.Println("Switch tanpa kondisi:")
	suhu := 30
	switch {
	case suhu >= 35:
		fmt.Println("Sangat panas")
	case suhu >= 30:
		fmt.Println("Panas")
	case suhu >= 25:
		fmt.Println("Hangat")
	default:
		fmt.Println("Sejuk")
	}
}