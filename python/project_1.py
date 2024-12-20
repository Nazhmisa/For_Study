while True:
    symbol = input("Введите операцию расчета (+,*,/ или -): ")

    if symbol == '+':
        number_1 = int(input("Введите первое число: "))
        number_2 = int(input("Введите второе число: "))
        result = number_1 + number_2
        print(result)
    elif symbol == '*':
        number_1 = int(input("Введите первое число: "))
        number_2 = int(input("Введите второе число: "))
        result = number_1 * number_2
        print(result)
    elif symbol == '/':
        number_1 = int(input("Введите первое число: "))
        number_2 = int(input("Введите второе число: "))
        result = number_1 / number_2
        print(result)
    elif symbol == '-':
        number_1 = int(input("Введите первое число: "))
        number_2 = int(input("Введите второе число: "))
        result = number_1 - number_2
        print(result)
    elif symbol == 'exit':
        print('Вы прекращаете работу в калькуляторе. Пока.')
        break
    else:
        print('Такой операции не предусмотрено, выберите другой знак из перечисленных')
        continue
