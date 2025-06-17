// Map (Карты) в Go
package main

import (
	"fmt"
)

// Структуры (Structs) в Go
type Person struct {
	Name  string
	Age   int
	Hight float64
}

func main() {
	// Создание карты
	// map[ключ]значение
	// map[тип ключа]тип значения
	// map[string]int
	m := make(map[string]int)

	// Добавление элементов в карту
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4

	// Вывод карты
	fmt.Println(m)

	// Получение значения по ключу
	fmt.Println(m["one"])

	// Проверка наличия ключа в карте
	value, ok := m["four"]
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key 'four' not found")
	}
	// Удаление элемента из карты
	delete(m, "two")
	fmt.Println("After deletion:", m)
	// Итерация по карте
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Проверка длины карты
	fmt.Println("Length of map:", len(m))
	// Создание карты с инициализацией
	n := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 7,
	}
	fmt.Println("Initialized map:", n)

	// Создание экземпляра структуры
	p := Person{
		Name:  "Alice",
		Age:   30,
		Hight: 1.75}
	// Вывод структуры
	fmt.Println("Person:", p)
	// Функция для вывода информации о человеке
	printPersonInfo(p)
}

// Функция для вывода информации о человеке
func printPersonInfo(person Person) {
	fmt.Printf("Name: %s\nAge: %d \nHeight: %.2f\n", person.Name, person.Age, person.Hight)
}
