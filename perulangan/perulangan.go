package perulangan

import "fmt"

func Perulangan() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//penggunaan dengan argumen hanya kondisi

	var y = 0

	for y < 5 {
		fmt.Println("angka", y)
		y++
	}

	//penggunaan for tanpa kondisi

	var j = 0
	for {
		fmt.Println("number", j)
		j++
		if j == 5 {
			break
		}
	}

	var xs = "123" //string
	for i, v := range xs {
		fmt.Println("index", i, "value", v)
	}
	var ys = [5]int{10, 20, 30, 40, 50} //array
	for _, v := range ys {
		fmt.Println("value", v)
	}
	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Println("value", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("key", k, "value", v)
	}
	// boleh juga k dan v nya diabaikan
	for range kvs {
		fmt.Println("done")

		//break and continue

		for z := 1; z <= 10; z++ {
			if z%2 == 1 {
				continue
			}
			if z > 8 {
				break
			}
			fmt.Println("angka", z)
		}
	}

	//perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println() // sama seperti \n
	}
	// outerLoop
outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
