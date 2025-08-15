package seleksi

import "fmt"

func Seleksi() {
	var value = 90

	if value < 50 {
		fmt.Println("mengulang", value)
	} else if value >= 50 && value < 70 {
		fmt.Println("remidi", value)
	} else if value >= 70 && value <= 100 {
		fmt.Println("lulus", value)
	} else {
		fmt.Println("salah memasukkan nilai!!!")
	}

	var point = 50

	switch point {
	case 10:
		fmt.Println("ngulang")
	case 20, 30, 40, 50:
		fmt.Println("lanjut")
	default:
		fmt.Println("masukkan angka 10-50")
	}
	// lihat perbedaan dibagian key switch, yang bawah menggunakan fullthrough
	switch {
	case point < 20:
		fmt.Println("ngulang")
	case point >= 20 && point <= 40:
		fmt.Println("cukup")
	case point > 40:
		fmt.Println("lulus")
	default:
		fmt.Println("masukkan angka 10-50!!!")
	}

}
