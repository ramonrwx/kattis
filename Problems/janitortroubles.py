import math


def max_area(sides: list) -> float:
    a, b, c, d = sides
    semi_perimeter = sum(sides) / 2
    return math.sqrt((semi_perimeter - a) *
                     (semi_perimeter - b) *
                     (semi_perimeter - c) *
                     (semi_perimeter - d))


sides = [int(x) for x in input().split()]
print(max_area(sides))
