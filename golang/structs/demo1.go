package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "Lenova"})
	// çalişmaz fmt.Println(product{"Bilgisayar", 500})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 500})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
