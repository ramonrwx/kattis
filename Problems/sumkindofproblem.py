cases = [list(map(int, input().split())) for _ in range(int(input()))]

for case in cases:
    K, N = case
    sqr = N*N

    print(K, (sqr + N) >> 1, sqr, sqr + N)
