package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 1000

	if cekilmekistenen > hesap {
		fmt.Println("Hesaptaki para yetersiz!")

	} else if cekilmekistenen == hesap {
		fmt.Println("Paranız hazırlanıyor")
		fmt.Println("Dikkat!! Hesapta paranız kalmadı")

	} else {
		fmt.Println("Paranız hazırlanıyor")
	}

}
