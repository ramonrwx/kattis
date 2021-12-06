a, b, c = input().split()

for i in range(1, int(c)+1):
    a, b, c = int(a), int(b), int(c)
    if i % a == 0 and i % b == 0:
        print('FizzBuzz')
    elif i % b == 0:
        print('Buzz')
    elif i % a == 0:
        print('Fizz')
    else:
        print(i)
