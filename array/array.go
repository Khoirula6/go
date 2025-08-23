package array

import "fmt"

func Array() {
	var names [4]string
	names[0] = "Mbak"
	names[1] = "Cantik"
	names[2] = "khoirul"
	names[3] = "anam"

	fmt.Println(names)
	fmt.Println(names[0], names[1], names[2], names[3])
	// array horizontal
	var fruits = [3]string{"apple", "mango", "orange"}
	fmt.Println("jumlah elemen \t\t", len(fruits))
	fmt.Println("isi elemen \t", fruits)

	// array dengan gaya vertikal
	var fish = [3]string{
		"koi",
		"lele",
		"gurame",
	}

	fmt.Println(fish)

	//array tanpa jumlah elemennya
	var number = [...]int{1, 2, 6, 7}
	fmt.Println("data array ", number)
	fmt.Println("jumlah elemen ", len(number))

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{4, 5, 6}, [3]int{1, 2, 3}}
	fmt.Println("numbers1 ", numbers1)

	var numbers2 = [2][3]int{{1, 2, 3}, {7, 8, 9}}
	fmt.Println("numbers2", numbers2)

	var numbers3 = [2][3]int{
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(numbers3)

	//array menggunakan for
	var fruits1 = [4]string{"apel", "water melon", "orange", "banana"}

	for i := 0; i < len(fruits1); i++ {
		// fmt.Println(i, "fruits ", fruits1[i])
		fmt.Printf("elemen %d: %s \n", i, fruits1[i])

	}

	//array menggunkan for range
	var animals = [4]string{"lion", "crocodile", "camel", "cow"}

	for _, animal := range animals {
		// for i, animal := range animals {

		fmt.Printf("hewan: %s\n", animal)
		// fmt.Println("hewan : ", i, animal)

	}

	// array menggunkan make
	var fishs = make([]string, 2)
	fishs[0] = "lele"
	fishs[1] = "emas"

	fmt.Println(fishs)

}
