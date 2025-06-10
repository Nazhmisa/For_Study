// Преобразование одного типа данных в другой
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Преобразование целого числа в строку
	num := 42
	strNum := strconv.Itoa(num)
	fmt.Printf("Целое число %d в строку: %s\n", num, strNum)

	// Преобразование строки в целое число
	str := "123"
	intNum, err := strconv.Atoi(str)
	// Обработка ошибки преобразования
	// Если строка не может быть преобразована в целое число, будет возвращена ошибка
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	} else {
		fmt.Printf("Строка %s в целое число: %d\n", str, intNum)
	}

	// Преобразование вещественного числа в строку
	floatNum := 3.14
	strFloat := strconv.FormatFloat(floatNum, 'f', -1, 64)
	fmt.Printf("Вещественное число %.2f в строку: %s\n", floatNum, strFloat)

	// Преобразование строки в вещественное число
	strFloatNum := "2.71828"
	floatValue, err := strconv.ParseFloat(strFloatNum, 64)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	} else {
		fmt.Printf("Строка %s в вещественное число: %.5f\n", strFloatNum, floatValue)
	}

	// Преобразование строки в число с плавающей запятой
	strFloatValue := "3.14"
	floatValue2, err := strconv.ParseFloat(strFloatValue, 64)
	fmt.Printf("Строка %s в число с плавающей запятой: %.2f\n", strFloatValue, floatValue2)

	// Константы в Go
	const AppName = "MyApp"
	const Version = 1.0
	const SecondsInHour = 3600
	fmt.Printf("Имя приложения: %s, Версия: %.1f, Секунд в часе: %d\n", AppName, Version, SecondsInHour)

	// Операторы в Go
	// В Go есть несколько типов операторов, которые можно использовать для выполнения различных операций. Вот основные из них:
	// 1. Арифметические операторы: +, -, *, /, %
	// 2. Операторы сравнения: ==, !=, <, >, <=, >=
	// 3. Логические операторы: &&, ||, !
	// 4. Побитовые операторы: &, |, ^, <<, >>
	// 5. Операторы присваивания: =, +=, -=, *=, /=, %=
	// 6. Операторы инкремента и декремента: ++, --
	// 7. Операторы управления потоком: if, else, switch, for, break, continue
	// 8. Операторы управления горутинами: go, select, defer
	// 9. Операторы управления каналами: <- (для отправки и получения данных из каналов)

	var (
		x = 15
		y = 4
	)

	// Арифметические операторы
	sum := x + y       // Сложение
	diff := x - y      // Вычитание
	product := x * y   // Умножение
	quotient := x / y  // Деление
	remainder := x % y // Остаток от деления
	fmt.Printf("Сумма: %d, Разность: %d, Произведение: %d, Частное: %d, Остаток: %d\n", sum, diff, product, quotient, remainder)

	// Операторы сравнения
	isEqual := x == y          // Равенство
	isNotEqual := x != y       // Неравенство
	isLess := x < y            // Меньше
	isGreater := x > y         // Больше
	isLessOrEqual := x <= y    // Меньше или равно
	isGreaterOrEqual := x >= y // Больше или равно
	fmt.Printf("Равенство: %t, Неравенство: %t, Меньше: %t, Больше: %t, Меньше или равно: %t, Больше или равно: %t\n", isEqual, isNotEqual, isLess, isGreater, isLessOrEqual, isGreaterOrEqual)
	// Логические операторы
	isTrue := true
	isFalse := false
	logicalAnd := isTrue && isFalse // Логическое И
	logicalOr := isTrue || isFalse  // Логическое ИЛИ
	logicalNot := !isTrue           // Логическое НЕ
	fmt.Printf("Логическое И: %t, Логическое ИЛИ: %t, Логическое НЕ: %t\n", logicalAnd, logicalOr, logicalNot)

}
