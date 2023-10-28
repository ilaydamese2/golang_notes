package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "Kırşehir"
	sehirler[2] = "İzmir"
	sehirler[3] = "İstanbul"
	sehirler[4] = "Kırıkkale"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}

}
