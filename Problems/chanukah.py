cases = int(input())

for i in range(cases):
    data, days = [int(x) for x in input().split(' ')]

    days = days + (days*(days+1)) / 2
    print(data, int(days))
