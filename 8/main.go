package main

import (
	"fmt"
	"math"
)

// 8. Введите число с плавающей точкой и преобразуйте его в целое число (округлить).
// Ввод: 3.14 → 3
// Ввод: 7.99 → 8

func main() {
	var floatNum float64 = 5.43
	fmt.Println(math.Round(float64(floatNum)))
}
