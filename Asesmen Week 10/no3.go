package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0

	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}

	return maxIdx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}

	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan ===")

	for i := 0; i < nProv; i++ {

		if tumbuh[i] > 0.02 {

			prediksi := int((1 + tumbuh[i]) * float64(pop[i]))

			fmt.Println(prov[i], ":", prediksi)
		}
	}
}

func main() {

	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string

	fmt.Println("=== Masukkan Data Provinsi ===")

	InputData(&prov, &pop, &tumbuh)

	fmt.Print("\nMasukkan nama provinsi yang dicari : ")
	fmt.Scan(&cari)

	idxCepat := ProvinsiTercepat(tumbuh)

	fmt.Println("\nProvinsi dengan pertumbuhan tercepat :", prov[idxCepat])

	idxCari := IndeksProvinsi(prov, cari)

	fmt.Println("Indeks provinsi :", idxCari)

	Prediksi(prov, pop, tumbuh)
}
