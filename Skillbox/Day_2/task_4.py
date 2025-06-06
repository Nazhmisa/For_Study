# str_of_palindrom = str(input("Введите палиндром: "))

# print(str_of_palindrom)

# n = 0
# yo = 0
# for i in str_of_palindrom:
#     n += 1
#     if i == str_of_palindrom[-n]:
#         print("yo")
#         if "yo":
#             yo += 1
#     elif i == " ":
#         continue
#     else:
#         print("ne yo")
# if yo == len(str_of_palindrom):
#     print("Это палиндром")
# else:
#     print("Не палиндром")
# print(n)

# Ввод строки
str_of_palindrom = input("Введите строку: ")

# Удаляем пробелы и приводим к одному регистру для корректного сравнения
cleaned_str = str_of_palindrom.replace(" ", "").lower()

# Проверка на палиндром
if cleaned_str == cleaned_str[::-1]:  # Сравниваем строку с её перевёрнутым вариантом
    print("Да, это палиндром!")
else:
    print("Нет, это не палиндром!")


# Объяснение:
# Очистка строки:

# Удаляем пробелы (replace(" ", "")), чтобы строки вроде "а роза упала на лапу Азора" распознавались как палиндромы.
# Приводим к нижнему регистру (lower()), чтобы регистр букв не влиял на сравнение.
# Проверка палиндрома:

# Используем синтаксис срезов [::-1], чтобы перевернуть строку.
# Сравниваем очищенную строку с её перевёрнутой копией.
# Простой вывод:

# Если строки совпадают, это палиндром, иначе — нет.