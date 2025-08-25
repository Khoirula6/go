package slice

import "fmt"

// perbedaan pembuatan array dan slice. slice tidak perlu mendefinisikan jumlah elemen
// var fruitsA = []string{"apple", "grape"} // slice
// var fruitsB = [2]string{"banana", "melon"} // array
// var fruitsC = [...]string{"papaya", "grape"} // array

func Slice() {
	var fruit = []string{"banana", "orange", "manggo"}
	fmt.Println(fruit[1])
	fmt.Println(fruit[0:3]) //dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-3.

	// atau
	var newFruit = fruit[0:2] //dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2.
	fmt.Println(newFruit)

	//len = menghitung jumlah elemen slice
	var food = []string{"martabak", "kebab", "roti", "sosis"}
	fmt.Println(len(food))

	//append untuk menambahkan elemen slice
	var foods = append(food, "lontong")
	fmt.Println(len(foods))

	fmt.Println(food)
	fmt.Println(foods)

}
