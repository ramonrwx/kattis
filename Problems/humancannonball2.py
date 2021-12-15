import math

cases = int(input())

for _ in range(cases):
    v, angle, x1, h1, h2 = (float(x) for x in input().split())
    angle = angle / 180 * math.pi
    h1 = h1 + 1
    h2 = h2 - 1

    t = x1 / (math.cos(angle) * v)
    y = v * t * math.sin(angle) - 0.5 * 9.81 * t * t
    print('Safe') if y >= h1 and y <= h2 else print('Not Safe')
