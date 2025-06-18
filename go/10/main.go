// Методы в Go

package main

import (
	"fmt"
)

// Структура с методом
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод для вычисления площади
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Метод для вычисления периметра
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func main() {
	// Создание экземпляра структуры Rectangle
	rect := Rectangle{Width: 5, Height: 3}

	// Вызов методов
	area := rect.Area()
	perimeter := rect.Perimeter()

	// Вывод результатов
	fmt.Printf("Area of rectangle: %.2f\n", area)
	fmt.Printf("Perimeter of rectangle: %.2f\n", perimeter)
}
