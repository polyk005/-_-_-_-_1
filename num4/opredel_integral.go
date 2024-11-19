package main

import (
	"fmt"
	"math"
)

// Определяем функцию, которую нужно интегрировать
func f(x float64) float64 {
	return (8 * math.Exp(x) / 20) + math.Pow(x, 3) - 3*math.Pow(x, 2) + (8 * x / 10) + 2
}

// Метод трапеций
func trap(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	totalSum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		totalSum += f(x)
	}
	return totalSum * h
}

func main() {
	a := 0.0
	b := 2.0

	// Вычисляем интеграл с использованием метода трапеций
	integral := trap(a, b, 1000) // Используем 1000 подинтервалов для большей точности

	// Оценки ошибок
	msp4 := math.Abs(integral - 1.03229855)
	msp8 := math.Abs(integral - 1.03724482)
	mspu := math.Abs(integral - 1.03889358)

	msi4 := math.Abs(integral - 1.03912102)
	msi8 := math.Abs(integral - 1.03891937)
	msiu := math.Abs(integral - 1.03890593)

	mtr4 := math.Abs(integral - 1.05216101)
	mtr8 := math.Abs(integral - 1.04222978)
	mtru := math.Abs(integral - 1.03891937)

	// Выводим результат
	fmt.Printf("Определенный интеграл от %.2f до %.2f равен %.6f\n", a, b, integral)
	fmt.Printf("msp4 равен %.6f\n", msp4)
	fmt.Printf("msp8 равен %.6f\n", msp8)
	fmt.Printf("mspu равен %.6f\n", mspu)
	fmt.Printf("msi4 равен %.6f\n", msi4)
	fmt.Printf("msi8 равен %.6f\n", msi8)
	fmt.Printf("msiu равен %.6f\n", msiu)
	fmt.Printf("mtr4 равен %.6f\n", mtr4)
	fmt.Printf("mtr8 равен %.6f\n", mtr8)
	fmt.Printf("mtru равен %.6f\n", mtru)
}
