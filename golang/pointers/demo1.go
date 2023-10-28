package pointers

import "fmt"

func Demo1(sayi *int) { // * adres anlamına gelir
	*sayi = *sayi + 1

	fmt.Println("Demodaki sayi : ", *sayi)

}
