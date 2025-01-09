#num = int(input("Введите количество карточек: "))
#cards = []
#sum_of_cards = 0
#sum_of_num = 0
#for x in range(num-1):
#    num_of_card = int(input("Введите номер оставшейся карточки: "))
#    cards.append(num_of_card)
#    if x in cards:
#        sum_of_cards += x
#for i in range(num):
#    sum_of_num += i
#lost_card = sum_of_num - sum_of_cards
#print(f"Номер потерянной карты: {lost_card}")


num = int(input("Введите количество карточек: "))
sum_of_cards = 0

# Суммируем оставшиеся карточки
for _ in range(num - 1):
    num_of_card = int(input("Введите номер оставшейся карточки: "))
    sum_of_cards += num_of_card

# Находим сумму всех карточек от 1 до N
sum_of_num = num * (num + 1) // 2

# Ищем номер потерянной карточки
lost_card = sum_of_num - sum_of_cards
print(f"Номер потерянной карточки: {lost_card}")
