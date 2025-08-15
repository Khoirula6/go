package operator

import "fmt"

func Operator() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t)\n", value, isEqual)
	// %t for bool  %d for int

	var cantik = true
	var jelek = false

	var leftAndRightn = cantik && jelek
	fmt.Printf("cantik && jelek = \t(%t)\n", leftAndRightn) // kanan dan kiri apakah sama

	var leftOrRight = cantik || jelek
	fmt.Printf("cantik || jelek = \t(%t)\n", leftOrRight) // kanan atau kiri berbeda

	var leftReverse = !cantik
	fmt.Printf("!cantik \t\t(%t)\n", leftReverse)
	// \t adalah jarak (tab) untuk ditampilkan di terminal

}
