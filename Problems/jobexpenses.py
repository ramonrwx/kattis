input()
numbers = [x for x in input().split()]

total = 0
for number in numbers:
    if number[0] == '-':
        total += int(number)

print(-total)
