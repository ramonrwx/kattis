from math import hypot

n, w, h = (int(x) for x in input().split())


def fits_in_box(nums: list[int]) -> None:
    longest = hypot(w, h)
    for num in nums:
        print('DA') if num <= longest else print('NE')


fits_in_box(int(input()) for _ in range(n))
