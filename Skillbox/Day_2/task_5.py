
nums = int(input("Введите число: "))
new_num = 1

for num in range(nums):
    space_count = nums - num - 1
    print("  " * space_count, end="")
    for j in range(num+1):
        print(new_num, end="  ")
        new_num += 2
    print()