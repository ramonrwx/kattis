from math import hypot

n, w, h = (int(x) for x in input().split())
LONGEST = hypot(w, h)

def fits_in_box(num: int) -> None:
    print('DA') if num <= LONGEST else print('NE')

for _ in range(n):
    fits_in_box(int(input()))
