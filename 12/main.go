// 12. Сравнение двух входных чисел
// Проверьте: Если a > b — "a больше" Если a == b — "a равно b" Если a < b — "b больше
//  Найдите максимальное из двух входных чисел.
//  Найдите минимальное из трех входных чисел.
//  Проверьте, равны ли два входных числа.

package main

import "fmt"

func main() {
	var a int = 49
	var b int = 45

	switch {
	case a > b:
		fmt.Println("a bigger")
	case a == b:
		fmt.Println("a equel b")
	case a < b:
		fmt.Println("b bigger")
	default:
		fmt.Println("another case")
	}

}
