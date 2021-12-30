cases = int(input())
numbers = [int(x) for x in input().split()]

tmp = counter = total = 0
for i in range(cases):
    tmp = numbers[i]
    if tmp < 0:
        continue

    counter += 1
    total += tmp

print(f'{(total/counter):.10f}')
