# X = int(input("Введите количество мальчиков: "))
# Y = int(input("Введите количество девочек: "))

# str_of_bg = ""

# for x in range(X):
#     str_of_bg += "B"
#     for y in range(Y):
#         if "B" in str_of_bg:
#             str_of_bg += "G"
#             continue
# print(str_of_bg)

X = int(input("Введите количество мальчиков: "))
Y = int(input("Введите количество девочек: "))

str_of_bg = ""
if X > Y * 2 or Y > X * 2:
    print("Нет решения")
else:
    while X > 0 or Y > 0:
        if X > Y:
            str_of_bg += "B"
            X -= 1
            if Y > 0:
                str_of_bg += "G"
                Y -= 1
        elif Y > X:
            str_of_bg += "G"
            Y -= 1
            if X > 0:
                str_of_bg += "B"
                X -= 1
        else:
            str_of_bg += "B"
            X -= 1
            str_of_bg += "G"
            Y -= 1
    print(str_of_bg)