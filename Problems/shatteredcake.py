n = int(input())
p = int(input())

total = 0
for i in range(p-1, -1, -1):
    w, h = (int(x) for x in input().split())
    total += h*w

print(int(total/n))
