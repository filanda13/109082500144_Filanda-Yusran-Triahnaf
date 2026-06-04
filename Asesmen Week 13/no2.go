package main

import (
	"fmt"
)

type Pemain struct {
	Nama   string
	Gol    int
	Assist int
}

func main() {
	var n int
	fmt.Println("Masukkan Data Input : ")
	fmt.Scanln(&n)

	var daftarPemain [1001]Pemain

	for i := 0; i < n; i++ {
		var namaDepan, namaBelakang string
		var gol, assist int

		fmt.Scanf("%s %s %d %d\n", &namaDepan, &namaBelakang, &gol, &assist)

		daftarPemain[i] = Pemain{
			Nama:   namaDepan + " " + namaBelakang,
			Gol:    gol,
			Assist: assist,
		}
	}

	for i := 1; i < n; i++ {
		key := daftarPemain[i]
		j := i - 1

		for j >= 0 && (daftarPemain[j].Gol < key.Gol || (daftarPemain[j].Gol == key.Gol && daftarPemain[j].Assist < key.Assist)) {
			daftarPemain[j+1] = daftarPemain[j]
			j--
		}
		daftarPemain[j+1] = key
	}

	fmt.Println("\n\nHasil Sorting : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", daftarPemain[i].Nama, daftarPemain[i].Gol, daftarPemain[i].Assist)
	}
}
