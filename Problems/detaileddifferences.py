cases = int(input())

for _ in range(cases):
    chars_one = input()
    chars_two = input()

    print(chars_one)
    print(chars_two)

    for char, char2 in zip(chars_one, chars_two):
        print('.' if char == char2 else '*', end='')
    print('\n')
