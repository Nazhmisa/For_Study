# list_of_dishes = input("Введите названия блюд, разделённых точкой с запятой: ")
# # Разделяем строку на список блюд
# changed_list = list_of_dishes.split(";")
# # Создаём новый список для отформатированных названий
# formatted_dishes = []

# for word in changed_list:
#     # Убираем лишние пробелы и форматируем каждое название
#     changed_word = word.strip().title()
#     formatted_dishes.append(changed_word)

# # Объединяем отформатированные названия в строку
# new_list = ", ".join(formatted_dishes)

# # Выводим результат
# print(f"На данный момент в меню есть: {new_list}")


def nice_view(text):
    text = ", ".join(text.split(';'))
    return text.title()
site_menu = input('Введите доступное меню: ')
print('Доступное меню:', site_menu)
print('На данный момент в меню есть:', nice_view(site_menu))