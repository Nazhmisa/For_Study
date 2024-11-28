package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//Декларация и инициализация переменной

	var height int = 163

	fmt.Println("My height is:", height)

	//В чем полустрогость типизации:

	var count = 154
	fmt.Println("Count of houses is = ", count)

	//Декларация и инициализация переменных одного типа (множественный случай)

	var First, Second int = 12, 41
	fmt.Printf("В деревне %d кур и %d слонов\n", First, Second)

	//Декларирование блока переменных
	var (
		nameOfcharacter   string = "Carl"
		powerOfcharacter  string = "FireBall"
		HealthOfCharacter int    = 86
	)
	fmt.Printf("Статус персонажа: \nИмя персонажа: %s \nОсновной навык персонажа: %s \n Очки жизни персонажа: %d \n", nameOfcharacter, powerOfcharacter, HealthOfCharacter)
}
