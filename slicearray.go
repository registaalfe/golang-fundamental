package main

import "fmt"

func main() {
	birthMonth := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	slice := birthMonth[2:5] // Slicing the array from index 2 to 4

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println("Panjang slice:", len(slice))      // 3 → Maret, April, Mei
	fmt.Println("Kapasitas slice:", cap(slice)) 
}