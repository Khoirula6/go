package tipedata

import "fmt"

func PrintTipeData() {

	// tipe data numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -123456789

	fmt.Printf("bilangan positif = %d\n", positiveNumber)
	fmt.Printf("bilangan negatif = %d \n", negativeNumber)
	// %d = int menjadi string

	// tipe data desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal = %f\n", decimalNumber)
	fmt.Printf("bilangan desimal 3 angka dibelakang koma = %.3f \n", decimalNumber)

	//%f = float menjadi string

	// tipe data bool (boolean) = true / false
	var exist bool = true
	fmt.Printf("exist %t \n", exist)

	// %t = untuk memanggil bool

	//tipe data string
	var message string = "hello"
	fmt.Printf("message = %s\n", message)

	var pesan = `Nama saya "Mbak Cantik".
	salam kenal.
	mari belajar "Golang".`

	fmt.Println(pesan)

}
