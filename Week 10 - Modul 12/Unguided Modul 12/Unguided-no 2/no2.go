package main

import (
	"fmt"
)

func main() {
	var hitungSuara [21]int
	var suaraMasuk int = 0
	var suaraSah int = 0

	for {
		var angka int
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}

		suaraMasuk++

		if angka >= 1 && angka <= 20 {
			suaraSah++
			hitungSuara[angka]++
		}
	}

	idKetua := -1
	maksSuaraKetua := -1

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > maksSuaraKetua {
			maksSuaraKetua = hitungSuara[i]
			idKetua = i
		}
	}

	idWakil := -1
	maksSuaraWakil := -1

	for i := 1; i <= 20; i++ {
		if i == idKetua {
			continue
		}

		if hitungSuara[i] > maksSuaraWakil {
			maksSuaraWakil = hitungSuara[i]
			idWakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", idKetua)
	fmt.Printf("Wakil ketua: %d\n", idWakil)
}
