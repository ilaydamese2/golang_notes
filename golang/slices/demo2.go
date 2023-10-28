//slice kopyalama

package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Ankara", "Kırşehir", "İzmir"}
	fmt.Println(sehirler)
	sehirlerkopya := make([]string, len(sehirler))
	fmt.Println(sehirlerkopya)
	copy(sehirlerkopya, sehirler)
	fmt.Println(sehirlerkopya)
	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerkopya)

	fmt.Println(sehirler[1:3]) //1. indisten itibaren 3'e kadar al
	fmt.Println(sehirler[1:])  //1. indisten itibaren sona kadar aalır.
	fmt.Println(sehirler[:2])  //en baştan alır

}
