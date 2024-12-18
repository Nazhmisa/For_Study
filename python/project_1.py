symbol = input("Введите операцию расчета (+,*,/ или -): ")
number_1 = int(input("Введите первое число: "))
number_2 = int(input("Введите второе число: "))

if symbol == '+':
    result = number_1 + number_2
    print(result)
elif symbol == '*':
    result = number_1 * number_2
    print(result)
elif symbol == '/':
    result = number_1 / number_2
    print(result)
else:
    result = number_1 - number_2
    print(result)