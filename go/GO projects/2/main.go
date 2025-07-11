// Добавление проверки дублирования email при создании пользователя
package main

import (
	"encoding/json" // Для работы с JSON (кодирование/декодирование)
	"fmt"           // Для форматированного вывода
	"net/http"      // Для работы с HTTP (создание сервера, обработка запрос
	"strings"       // Для работы со строками (например, для приведения email к нижнему регистру)
)

type User struct {
	ID    int    `json:"id"`    // Поле ID с тегом для JSON
	Name  string `json:"name"`  // Поле имени с тегом
	Email string `json:"email"` // Поле email с тегом
}

var users []User  // Глобальная переменная для хранения списка пользователей
var idCounter int // Счетчик для генерации ID новых пользователей
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения списка пользователей
	if r.Method != http.MethodGet {
		// Если метод запроса не GET, возвращаем ошибку 405
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок ответа
	json.NewEncoder(w).Encode(users)                   // Кодируем список пользователей в JSON и отправляем
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для создания нового пользователя
	if r.Method != http.MethodPost {
		// Если метод запроса не POST, возвращаем ошибку 405
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newUser User // Создаем переменную для нового пользователя
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		// Если произошла ошибка при декодировании JSON, возвращаем ошибку 400
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" || newUser.Email == "" {
		// Проверяем, заполнены ли обязательные поля
		http.Error(w, "Имя и email обязательны", http.StatusBadRequest)
		return
	}

	// Проверяем, существует ли уже пользователь с таким email (без учета регистра)
	for _, user := range users {
		if strings.EqualFold(user.Email, newUser.Email) {
			http.Error(w, "Пользователь с таким email уже существует", http.StatusConflict)
			return
		}
	}

	idCounter++                    // Увеличиваем счетчик ID
	newUser.ID = idCounter         // Присваиваем новый ID пользователю
	users = append(users, newUser) // Добавляем нового пользователя в список

	w.WriteHeader(http.StatusCreated)  // Устанавливаем статус 201 Created
	json.NewEncoder(w).Encode(newUser) // Кодируем нового пользователя в JSON и отправляем
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения списка пользователей или создания нового пользователя
	switch r.Method {
	case http.MethodGet:
		getUsersHandler(w, r) // Если GET, вызываем обработчик для получения списка пользователей
	case http.MethodPost:
		createUserHandler(w, r) // Если POST, вызываем обработчик для создания нового пользователя
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed) // Для других методов возвращаем ошибку 405
	}
}
func main() {
	// Основная функция для запуска HTTP-сервера
	http.HandleFunc("/users", usersHandler)     // Устанавливаем обработчик для пути /users
	fmt.Println("Сервер запущен на порту 8080") // Выводим сообщение о запуске сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Запускаем сервер и обрабатываем возможные ошибки
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}

// Примечания к проекту:
// - Этот проект расширяет предыдущий пример, добавляя проверку на уникальность email при создании пользователя.
// - Он использует структуру `User` для представления пользователя с полями ID, Name и Email.
// - При создании пользователя проверяется, существует ли уже пользователь с таким email (без учета регистра).
// - Если пользователь с таким email уже существует, возвращается ошибка 409 Conflict.
