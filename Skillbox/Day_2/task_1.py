def binary_search(arr, target):
    # Указываем начальные границы диапазона
    left = 0
    right = len(arr) - 1

    while left <= right:
        # Вычисляем середину диапазона
        mid = (left + right) // 2

        # Проверяем число в середине
        if arr[mid] == target:
            return mid  # Возвращаем индекс найденного числа
        elif arr[mid] < target:
            left = mid + 1  # Сужаем диапазон поиска вправо
        else:
            right = mid - 1  # Сужаем диапазон поиска влево

    return -1  # Если число не найдено

# Пример использования
numbers = list(range(1, 1001))  # Список от 1 до 1000
target = int(input("Введите число для поиска (от 1 до 1000): "))
result = binary_search(numbers, target)

if result != -1:
    print(f"Число {target} найдено на индексе {result}.")
else:
    print(f"Число {target} не найдено.")