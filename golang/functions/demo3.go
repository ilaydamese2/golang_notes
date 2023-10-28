// Variadic fonksiyon

package functions

func ToplaVariadic(sayilar ...int) int { //sayilari dizi olarak ele alıyor
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}

	return toplam

}
