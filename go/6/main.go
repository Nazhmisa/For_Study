package main

import (
	"fmt"
)

func main() {
	// Управляющие конструкции (if/else)
	score := 40
	fmt.Println("Оценка:")
	if score >= 90 {
		fmt.Println("Отлично!")
	} else if score >= 75 && score < 90 {
		fmt.Println("Хорошо!")
	} else if score >= 60 && score < 75 {
		fmt.Println("Удовлетворительно!")
	} else {
		fmt.Println("Неудовлетворительно!")
	}
	// Циклы (for)
	fmt.Println("Цикл for:")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// Цикл с условием (for с условием)
	fmt.Println("Цикл for с условием (только четные числа):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	// Switch-case
	fmt.Println("Switch-case:")
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Отлично!")
	case "B":
		fmt.Println("Хорошо!")
	case "C":
		fmt.Println("Удовлетворительно!")
	case "D":
		fmt.Println("Неудовлетворительно!")
	default:
		fmt.Println("Неизвестная оценка!")
	}
}
