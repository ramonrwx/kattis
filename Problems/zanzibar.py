cases = int(input())

for _ in range(cases):
    values = [int(v) for v in input().split()]
    total = 0
    for i in range(1, len(values)):
        curr = values[i]
        prev = values[i-1]
        if curr/prev < 2:
            continue
        total += curr - prev * 2
    print(total)
