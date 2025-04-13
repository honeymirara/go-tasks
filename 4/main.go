// 4. Проверьте, является ли число четным или нечетным.
package main

import "fmt"

func main() {

	var value int = 100

	if value%2 == 0 {
		fmt.Println("Countable")
	} else {
		fmt.Println("Uncountable")
	}
}
