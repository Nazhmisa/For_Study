// Подключение реальной базы данных на SQLite

package main

import (
	"database/sql"  // Для работы с базой данных
	"encoding/json" // Для работы с JSON (кодирование/декодирование)

	// Для форматированного вывода
	"log"      // Для логирования ошибок
	"net/http" // Для работы с HTTP (создание сервера, обработка запросов)

	// Для работы с SQLite
	_ "modernc.org/sqlite"
	// Импортируем драйвер SQLite для работы с базой данных
)

var db *sql.DB // Глобальная переменная для базы данных

type User struct {
	// Структура для хранения информации о пользователе
	ID    int    `json:"id"`    // ID пользователя
	Name  string `json:"name"`  // Имя пользователя
	Email string `json:"email"` // Email пользователя
}

func initDB() {
	// Функция для инициализации базы данных
	var err error
	db, err = sql.Open("sqlite", "./users.db") // Открываем базу данных SQLite
	if err != nil {
		log.Fatal(err) // Логируем ошибку, если не удалось открыть базу данных
	}

	// Создаем таблицу пользователей, если она не существует
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);
	`
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	log.Println("База данных и таблица пользователей успешно инициализированы")
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для создания нового пользователя
	if r.Method != http.MethodPost {
		// Если метод запроса не POST, возвращаем ошибку 405
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newUser User                                // Создаем переменную для хранения пользователя
	err := json.NewDecoder(r.Body).Decode(&newUser) // Декодируем JSON из тела запроса
	if err != nil {
		// Если произошла ошибка при декодировании JSON, возвращаем ошибку 400
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", newUser.Name, newUser.Email)
	if err != nil {
		http.Error(w, "Email уже существует", http.StatusBadRequest)
		return
	} // Выполняем запрос на вставку нового пользователя в базу данных

	id, _ := result.LastInsertId()                                                 // Получаем ID нового пользователя
	newUser.ID = int(id)                                                           // Устанавливаем ID в структуру пользователя
	w.Header().Set("Content-Type", "application/json")                             // Устанавливаем заголовок ответа
	json.NewEncoder(w).Encode(newUser)                                             // Кодируем нового пользователя в JSON и отправляем в ответ
	log.Printf("Пользователь %s с ID %d успешно создан", newUser.Name, newUser.ID) // Логируем успешное создание пользователя
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения списка пользователей
	if r.Method != http.MethodGet {
		// Если метод запроса не GET, возвращаем ошибку 405
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name, email FROM users") // Выполняем запрос к базе данных
	if err != nil {
		// Если произошла ошибка при выполнении запроса, возвращаем ошибку 500
		http.Error(w, "Ошибка при получении пользователей", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Закрываем rows после использования

	var users []User // Создаем срез для хранения пользователей
	for rows.Next() {
		var user User // Создаем переменную для хранения пользователя
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Ошибка при чтении данных", http.StatusInternalServerError)
			return
		}
		users = append(users, user) // Добавляем пользователя в срез
	}

	w.Header().Set("Content-Type", "application/json")  // Устанавливаем заголовок ответа
	json.NewEncoder(w).Encode(users)                    // Кодируем список пользователей в JSON и отправляем в ответ
	log.Println("Список пользователей успешно получен") // Логируем успешное получение списка пользователей
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Универсальный обработчик для создания и получения пользователей
	switch r.Method {
	case http.MethodPost:
		createUsersHandler(w, r) // Если метод POST, вызываем обработчик создания пользователя
	case http.MethodGet:
		getUsersHandler(w, r) // Если метод GET, вызываем обработчик получения списка пользователей
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed) // Если метод не поддерживается, возвращаем ошибку 405
	}
}

func main() {
	// Основная функция, где запускается сервер
	initDB() // Инициализируем базу данных

	http.HandleFunc("/users", usersHandler) // Обработчик для создания пользователей

	log.Println("Сервер запущен на порту 8080") // Логируем запуск сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		defer db.Close() // Закрываем базу данных при завершении работы
		// Если произошла ошибка при запуске сервера, логируем ее
		log.Fatal(err) // Логируем ошибку, если сервер не удалось запустить
	}
}
