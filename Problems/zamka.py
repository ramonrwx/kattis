L = int(input())
D = int(input())
X = int(input())

num_max = 0
num_min = 0


def sum_digits(n):
    return sum(int(x) for x in str(n))


for i in range(L, D+1):
    if (sum_digits(i) == X):
        num_min = i
        break

for j in range(D, L-1, -1):
    if (sum_digits(j) == X):
        num_max = j
        break

print(num_min)
print(num_max)
