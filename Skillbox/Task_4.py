# Получаем входное время
hour = int(input("Введите время (в часах от 0 до 23): "))

# Проверяем, находится ли время в запрещённых интервалах
is_open = hour in range(8, 22) and hour not in range(10, 12) and hour not in range(14, 15) and hour not in range(18, 20)

# Выводим результат
print("Можно получить посылку" if is_open else "Посылку получить нельзя")