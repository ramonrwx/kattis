p, n = (int(x) for x in input().split())
parts = set()

for i in range(n):
    parts.add(input())
    if len(parts) == p:
        print(i + 1)
        break
else:
    print('paradox avoided')
