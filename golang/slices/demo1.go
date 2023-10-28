package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "İlayda"
	isimler[1] = "Emo"
	isimler[2] = "kaan"
	isimler = append(isimler, "Tuğçe") //append yeni slice yeni array oluşturur

	fmt.Println(isimler)
	fmt.Println(len(isimler)) // len isimlerin kaç elemandan oluştuğunu gösteririr

}
