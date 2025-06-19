// Интерфейсы и их реализация
package main

import (
	"fmt"
)

// Интерфейс Animal
type Animal interface {
	Speak() string
}

// Интерфейс Animal определяет метод Speak, который должен быть реализован любым типом, который его использует
// Любой тип, который реализует метод Speak, может быть использован как Animal
// Это позволяет создавать функции, которые могут работать с различными типами животных, не зная их конкретных реализаций
// Это пример полиморфизма в Go, где разные типы могут быть использованы взаимозаменяемо, если они реализуют один и тот же интерфейс

// Интерфейс Shape
type Shape interface {
	area() float64 // Метод area, который должен быть реализован для вычисления площади фигуры
}

// Структура Dog
type Dog struct {
	Name string
}

// Реализация метода Speak для Dog
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name) // Возвращает строку с именем и звуком собаки, Sprintf используется для форматирования строки
}

// Структура Cat
type Cat struct {
	Name string
}

// Реализация метода Speak для Cat
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Meow!", c.Name) // Возвращает строку с именем и звуком кошки, Sprintf используется для форматирования строки
}

// Функция MakeSound принимает интерфейс Animal и вызывает метод Speak
func MakeSound(a Animal) {
	fmt.Println(a.Speak()) // Выводит результат вызова метода Speak для переданного животного
}

// Структура Circle
type Circle struct {
	Radius float64 // Радиус круга
}

// Реализация метода area для Circle
func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius // Вычисляет площадь круга по формуле πr²
}

// Структура Rectangle
type Rectangle struct {
	Width  float64 // Ширина прямоугольника
	Height float64 // Высота прямоугольника
}

// Реализация метода area для Rectangle
func (r Rectangle) area() float64 {
	return r.Width * r.Height // Вычисляет площадь прямоугольника по формуле ширина * высота
}

// Функция PrintArea принимает интерфейс Shape и выводит площадь фигуры
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.area()) // Выводит площадь фигуры, округленную до двух знаков после запятой
}

// Главная функция
func main() {
	dog := Dog{Name: "Buddy"}                   // Создание экземпляра Dog с именем Buddy
	cat := Cat{Name: "Whiskers"}                // Создание экземпляра Cat с именем Whiskers
	circle := Circle{Radius: 5}                 // Создание экземпляра Circle с радиусом 5
	rectangle := Rectangle{Width: 4, Height: 6} // Создание экземпляра Rectangle с шириной 4 и высотой 6
	PrintArea(circle)                           // Вызов функции PrintArea для круга
	PrintArea(rectangle)                        // Вызов функции PrintArea для прямоугольника
	MakeSound(dog)                              // Вызов функции MakeSound для собаки
	MakeSound(cat)                              // Вызов функции MakeSound для кошки
}
