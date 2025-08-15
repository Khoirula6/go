package variabel

import (
	"fmt"
)

func PrintVariabel() {
	//deklarasi variabel beserta tipe data
	var firstName string = "Khoirul"

	var lastName string
	lastName = "Anam"

	fmt.Println("Hello, my name: ", firstName, lastName+"!")
	fmt.Printf("hello, my name: %s %s !\n", firstName, lastName)

	//
	var animal = "kucing" //variabel tanpan tipedata

	animalName := "mbakku" // variabel tanpa var dan tanpa tipe data (hanya bisa digunakan dalam blok func)
	fmt.Println("hi", animal, animalName)

	var bird, cat, fish string // deklarasi multi variabel
	// bird = "merpati"
	// cat = "kucing oren"
	// fish = "ikan emas"
	bird, cat, fish = "merpati", "kucing oren", "ikan emas"

	fmt.Printf("peliharaanku ada %s %s dan %s \n", bird, cat, fish)

	var one, two, three = "satu", 2, "tiga"
	fmt.Println(one, two, three)

	_ = "kamu" //variabel underscore
	_ = "caca"
	hero, _ := "miya", "laila"
	fmt.Println(hero)

	color := new(string) // variabel color menampung data bertipe pointer string

	fmt.Println(color)  //menampilkan alamat memory
	fmt.Println(*color) // " "
}
