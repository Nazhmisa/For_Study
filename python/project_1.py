while True:
    symbol = input("Введите операцию расчета (+,*,/ или -): ")

    if symbol in ['+', '*', '-', '/']:
        number_1 = int(input("Введите первое число: "))
        number_2 = int(input("Введите второе число: "))
    
        if symbol == '+':
            result = number_1 + number_2
            print(f'Результат вычислений: {number_1} + {number_2} = {result}')
        elif symbol == '*':
            result = number_1 * number_2
            print(f'Результат вычислений: {number_1} * {number_2} = {result}')
        elif symbol == '/':
            if number_2 == 0:
                print('На ноль делить нельзя!')
                continue
            result = number_1 / number_2
            print(f'Результат вычислений: {number_1} / {number_2} = {result}')
        elif symbol == '-':
            result = number_1 - number_2
            print(f'Результат вычислений: {number_1} - {number_2} = {result}')
    elif symbol == 'exit':
        print('Вы прекращаете работу в калькуляторе. Пока.')
        break
    else:
        print('Такой операции не предусмотрено, выберите другой знак из перечисленных')
        continue
