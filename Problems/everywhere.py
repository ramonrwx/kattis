cases = int(input())

for _ in range(cases):
    cities = []
    cities_cases = int(input())
    for _ in range(cities_cases):
        cities.append(input())
    print(len(set(cities)))
