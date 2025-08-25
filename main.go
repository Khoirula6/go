package main

import (
	"fmt"
	"hello-world/array"
	"hello-world/operator"
	"hello-world/perulangan"
	"hello-world/seleksi"
	"hello-world/slice"
	tipedata "hello-world/tipe-data"
	"hello-world/variabel"
)

func main() {
	fmt.Println("Hello World")
	variabel.PrintVariabel()
	tipedata.PrintTipeData()
	operator.Operator()
	seleksi.Seleksi()
	perulangan.Perulangan()
	array.Array()
	slice.Slice()
}
