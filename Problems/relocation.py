_, n = (int(v) for v in input().split())
locations = [int(v) for v in input().split()]

for _ in range(n):
    a, b, c = (int(v) for v in input().split())
    if a != 1:
        print(abs(locations[b - 1] - locations[c - 1]))
    else:
        locations[b - 1] = c
