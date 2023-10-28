package loops

import "fmt"

//
func Demo3() {
	aklimdakisayi := 80
	tahminedilensayi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminedilensayi) //80

	for aklimdakisayi != tahminedilensayi {

		if tahminedilensayi < aklimdakisayi {
			fmt.Println("Daha büyük bir sayı giriniz!!")
			fmt.Scanln(&tahminedilensayi)
		}

		if tahminedilensayi > aklimdakisayi {
			fmt.Println("Daha küçük bir sayı giriniz!!")
			fmt.Scanln(&tahminedilensayi)
		}

	}

	if tahminedilensayi == aklimdakisayi {
		fmt.Println("Bildiniz..")
	}

}
