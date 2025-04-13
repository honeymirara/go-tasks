package main

//1. Проверьте возраст у входных данных :

// Если < 18 — "Вы несовершеннолетний".
// Если 18-60 — "Вы взрослый".
// Если > 60 — "Вы пожилой человек".

import "fmt"

func main() {
	var age int = 4
	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age > 18 && age < 60 {
		fmt.Println("Вы взрослый")
	} else if age > 60 {
		fmt.Println("Вы пожилой человек")
	}

}
