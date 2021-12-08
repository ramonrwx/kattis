def is_harshad(num):
    x, sum = num, 0
    while x > 0:
        sum += x % 10
        x = x // 10
    return num % sum == 0


number = int(input())

while 1:
    if is_harshad(number):
        print(number)
        break
    number += 1
