package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 900

	//Go dilinin derlenen halinde if'in parantezi kalkar

	if cekilmekistenen > hesap {
		fmt.Print("Hesaptaki para yetersiz!")
	}

	if cekilmekistenen <= hesap {
		fmt.Println("Paranız hazırlanıyor")
		hesap = hesap - cekilmekistenen
	}

	fmt.Println("Bitti. Hesaptaki para : " + fmt.Sprintf("%v", hesap))
}
