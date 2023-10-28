package maps

import "fmt"

func Demo1() {

	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı :", len(sozluk))
	delete(sozluk, "book") //eleman siler
	fmt.Println("Eleman sayısı :", len(sozluk))

	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)

}
