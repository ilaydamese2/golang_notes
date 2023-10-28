package arrays

import "fmt"

func Demo4() {

	//2 satır 3 sütundan oluşuyor dize
	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[0][2] = 5
	sayilar[1][0] = 2
	sayilar[1][1] = 4
	sayilar[1][2] = 6
	//mesela sayilar[2][0]=1 bunu kabul etmez

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")

		}
		fmt.Println()

	}

	//fmt.Println(sayilar[1][1])

}
