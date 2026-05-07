package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func nilaiPertama(arr arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(arr arrayMahasiswa, n int, nim string) int {
	max := -1

	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}

	return max
}

func main() {
	var arr arrayMahasiswa
	var n int
	var cariNIM string

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].nama, &arr[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cariNIM)

	pertama := nilaiPertama(arr, n, cariNIM)
	terbesar := nilaiTerbesar(arr, n, cariNIM)

	if pertama == -1 {
		fmt.Println("NIM tidak ditemukan")
	} else {
		fmt.Println("Nilai pertama dari NIM", cariNIM, "adalah", pertama)
		fmt.Println("Nilai terbesar dari NIM", cariNIM, "adalah", terbesar)
	}
}
