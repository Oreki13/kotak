package main

import (
	"fmt"
)

func main() {
	var data int
	fmt.Println("Masukan Jumlah Nilai :")
	fmt.Scan(&data)
	if data%2 == 0 {
		GambarGenap(data)
	} else {
		GambarGanjil(data)
	}
}

func GambarGanjil(data int) {
	tengah := data / 2
	var tmp string
	fmt.Println("---Panjang---")
	for i := 0; i < data; i++ {
		if i == tengah {
			for k := 0; k < data; k++ {
				tmp += "*"
			}
		} else {
			for w := 0; w < data; w++ {

				if w == 0 {
					tmp += "*"
				} else if w == data-1 {
					tmp += "*"
				} else {
					tmp += "="
				}
			}
		}
		fmt.Println(tmp)
		tmp = ""
	}
}

func GambarGenap(data int) {
	tengah := data / 2
	var tmp string
	fmt.Println("---Panjang---")
	for i := 0; i < data; i++ {
		if i == tengah || i == tengah-1 {
			for k := 0; k < data; k++ {
				tmp += "*"
			}
		} else {
			for w := 0; w < data; w++ {

				if w == 0 {
					tmp += "*"
				} else if w == data-1 {
					tmp += "*"
				} else {
					tmp += "="
				}
			}
		}
		fmt.Println(tmp)
		tmp = ""
	}
}
