// Композиция структур в Go

package main

import "fmt"

// Структура Person
type Person struct {
	Name string // Имя человека
	Age  int    // Возраст человека
}

// Структура Address
type Address struct {
	Street string // Улица адреса
	City   string // Город адреса
	Zip    string // Почтовый индекс адреса
}

// Структура Employee, которая включает в себя Person и Address
type Employee struct {
	Person           // Встраивание структуры Person
	Address          // Встраивание структуры Address
	Position string  // Должность сотрудника
	Salary   float64 // Зарплата сотрудника
}

// Функция PrintEmployeeInfo выводит информацию о сотруднике
func PrintEmployeeInfo(e Employee) {
	fmt.Printf("Name: %s\n", e.Name)         // Выводит имя сотрудника
	fmt.Printf("Age: %d\n", e.Age)           // Выводит возраст сотрудника
	fmt.Printf("Street: %s\n", e.Street)     // Выводит улицу адреса сотрудника
	fmt.Printf("City: %s\n", e.City)         // Выводит город адреса сотрудника
	fmt.Printf("Zip: %s\n", e.Zip)           // Выводит почтовый индекс адреса сотрудника
	fmt.Printf("Position: %s\n", e.Position) // Выводит должность сотрудника
	fmt.Printf("Salary: %.2f\n", e.Salary)   // Выводит зарплату сотрудника с двумя знаками после запятой
}

func main() {
	// Создаем экземпляр Employee с данными
	employee := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			Zip:    "12345",
		},
		Position: "Software Engineer",
		Salary:   75000.00,
	}

	// Выводим информацию о сотруднике
	PrintEmployeeInfo(employee)
}
