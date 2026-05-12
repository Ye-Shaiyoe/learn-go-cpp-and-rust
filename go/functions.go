// functions.go
// Penjelasan tentang fungsi dalam Go

package main

import (
	"fmt"
	"math"
)

// Fungsi dasar dengan parameter dan kembalian
func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

// Fungsi dengan kembalian bernama (named return values)
func hitungDiskon(harga, persentase float64) (diskon, hargaAkhir float64) {
	diskon = harga * persentase / 100
	hargaAkhir = harga - diskon
	return // akan mengembalikan diskon dan hargaAkhir
}

// Fungsi dengan kembalian variabel
func operasiMatematika(a, b float64) (tambah, kurang, kali, bagi float64) {
	tambah = a + b
	kurang = a - b
	kali = a * b
	if b != 0 {
		bagi = a / b
	} else {
		bagi = 0 // menghindari pembagian dengan nol
	}
	return
}

// Fungsi yang menerima jumlah parameter tidak tetap (variadic function)
func jumlahkan(angka ...float64) float64 {
	total := 0.0
	for _, nilai := range angka {
		total += nilai
	}
	return total
}

// Fungsi yang mengembalikan beberapa nilai termasuk error
func bagiDenganError(pembagi, penyebut float64) (float64, error) {
	if penyebut == 0 {
		return 0, fmt.Errorf("pembagian dengan nol tidak diperbolehkan")
	}
	return pembagi / penyebut, nil
}

func main() {
	fmt.Println("=== FUNGSI DALAM GO ===")

	// Memanggil fungsi dasar
	panjang, lebar := 10.5, 5.0
	luas := luasPersegiPanjang(panjang, lebar)
	fmt.Printf("Luas persegi panjang (%.1f x %.1f) = %.2f\n", panjang, lebar, luas)

	// Memanggil fungsi dengan kembalian bernama
	hargaAwal := 100000.0
	potongan := 15.0
	diskon, hargaAkhir := hitungDiskon(hargaAwal, potongan)
	fmt.Printf("Harga: %.0f, Diskon %0.0f%%, Potongan: %.0f, Harga akhir: %.0f\n",
		hargaAwal, potongan, diskon, hargaAkhir)

	// Memanggil fungsi dengan banyak kembalian
	a, b := 20.0, 4.0
	tambah, kurang, kali, bagi := operasiMatematika(a, b)
	fmt.Printf("\nOperasi pada %.1f dan %.1f:\n", a, b)
	fmt.Printf("Penjumlahan: %.1f\n", tambah)
	fmt.Printf("Pengurangan: %.1f\n", kurang)
	fmt.Printf("Perkalian: %.1f\n", kali)
	fmt.Printf("Pembagian: %.1f\n", bagi)

	// Memanggil fungsi variadic
	fmt.Println("\nFungsi variadic - jumlahkan:")
	fmt.Printf("Jumlah 1, 2, 3 = %.1f\n", jumlahkan(1, 2, 3))
	fmt.Printf("Jumlah 5.5, 4.5, 3.0, 2.0, 1.0 = %.1f\n", jumlahkan(5.5, 4.5, 3.0, 2.0, 1.0))
	fmt.Printf("Jumlah kosong = %.1f\n", jumlahkan()) // akan mengembalikan 0

	// Memanggil fungsi yang mengembalikan error
	fmt.Println("\nFungsi dengan error handling:")
	if hasil, err := bagiDenganError(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", hasil)
	}

	if hasil, err := bagiDenganError(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", hasil)
	}

	// Menggunakan fungsi dari package standar
	fmt.Println("\nMenggunakan fungsi dari package standar:")
	fmt.Println("Akar dari 16 =", math.Sqrt(16))
	fmt.Println("Pangkat 2^3 =", math.Pow(2, 3))
	fmt.Println("Nilai pi =", math.Pi)
}