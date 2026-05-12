// variables.go
// Penjelasan tentang variabel dalam Go

package main

import "fmt"

func main() {
	// Deklarasi variabel dengan var
	var nama string = "Budi"
	var umur int = 25

	// Deklarasi singkat dengan := (hanya di dalam fungsi)
	alamat := "Jakarta"

	// Multiple variabel
	var tinggi, berat float64 = 170.5, 65.2

	// Konstanta dengan const
	const phi float64 = 3.14
	const namaSekolah string = "SMK Teknologi"

	// Menampilkan nilai
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Tinggi:", tinggi, "cm")
	fmt.Println("Berat:", berat, "kg")
	fmt.Println("Phi:", phi)
	fmt.Println("Sekolah:", namaSekolah)

	// Tipe data default (zero value)
	var count int
	var namaKosong string
	fmt.Println("Zero value count:", count)
	fmt.Println("Zero value string:", namaKosong)
}