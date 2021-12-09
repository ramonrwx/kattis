cases = int(input())

for _ in range(cases):
    input()
    places = [int(x) for x in input().split()]
    print((max(places) - min(places))*2)
