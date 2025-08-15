package main

import (
	"fmt"
	"hello-world/operator"
	"hello-world/seleksi"
	tipedata "hello-world/tipe-data"
	"hello-world/variabel"
)

func main() {
	fmt.Println("Hello World")
	variabel.PrintVariabel()
	tipedata.PrintTipeData()
	operator.Operator()
	seleksi.Seleksi()
}
