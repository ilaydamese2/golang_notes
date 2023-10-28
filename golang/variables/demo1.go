package variables

//paket türü
import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya"
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	//açık değişken tanımlama
	// integer sayı tanımlama

	var kdv int = 15
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	// ondalıklı sayı tanımlama

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	//veri tipi bilinmiyorsa
	kdv3 := 25.2
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	//Mantıksal veri tipleri

	var durum bool //false true de olabilir ya true ya false

	var metin1 string = "İlayda"
	var metin2 string = "Emo"

	durum = metin1 != metin2
	fmt.Println(durum)
	fmt.Println(!durum)

}
