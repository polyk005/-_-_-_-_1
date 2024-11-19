package main

import (
	"fmt"
	"math"
)

const S = 8

//вычисляет функции
func f(x float64) float64 {
	return S*math.Exp(x)/20 + math.Pow(x, 3) - 3*math.Pow(x, 2) + S*x/10 + 2
}

// srPr метод средних прямоугольников
func srPr(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	totalSum := 0.0
	for i := 0; i < n; i++ {
		xSr := a + (float64(i)+0.5)*h
		totalSum += f(xSr)
	}
	return totalSum * h
}

func main() {
	a := 0.0
	b := 2.0
	n1 := 4
	integ4 := srPr(a, b, n1)
	fmt.Printf("С 4 частями: %.8f\n", integ4)

	n2 := 8
	integ8 := srPr(a, b, n2)
	fmt.Printf("С 8 частями: %.8f\n", integ8)

	if math.Abs(integ4-integ8) > 1e-8 {
		p := 2
		R := (integ8 - integ4) / (math.Pow(2, float64(p)) - 1)
		newInteg := integ8 + R
		fmt.Printf("Уточненное значение : %.8f\n", newInteg)
	}
}
