from math import pi, ceil, sin

h, v = (int(x) for x in input().split())
print(ceil(h / sin(pi * v / 180)))
