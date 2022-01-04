n, h, v = (int(v) for v in input().split())
print(max(h, n-h) * max(v, n-v) * 4)
