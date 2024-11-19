package main

import (
	"fmt"
	"math"
)

const S = 8

// значение функции
func f(x float64) float64 {
	return S*math.Exp(x)/20 + math.Pow(x, 3) - 3*math.Pow(x, 2) + S*x/10 + 2
}

// sim метод Симпсона
func sim(a, b float64, n int) float64 {
	if n%2 == 1 {
		panic("Количество подинтервалов должно быть четным.")
	}
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i += 2 {
		sum += 4 * f(a+float64(i)*h)
	}
	for i := 2; i < n-1; i += 2 {
		sum += 2 * f(a+float64(i)*h)
	}
	return sum * h / 3
}

func main() {
	a := 0.0
	b := 2.0
	n1 := 4
	integ4 := sim(a, b, n1)
	fmt.Printf("C 4 частями: %.8f\n", integ4)

	n2 := 8
	integ8 := sim(a, b, n2)
	fmt.Printf("C 8 частями: %.8f\n", integ8)

	if math.Abs(integ4-integ8) > 1e-8 {
		R := (integ8 - integ4) / (16 - 1)
		newInteg := integ8 + R
		fmt.Printf("Уточненное значение: %.8f\n", newInteg)
	}
}
