n, t = (int(x) for x in input().split())
numbers = [int(x) for x in input().split()]

if t == 1:
    print(7)
elif t == 2:
    if numbers[0] > numbers[1]:
        print('Bigger')
    elif numbers[0] == numbers[1]:
        print('Equal')
    else:
        print('Smaller')
elif t == 3:
    print(sorted(numbers[0:3])[1])
elif t == 4:
    print(sum(numbers))
elif t == 5:
    print(sum(x for x in numbers if x % 2 == 0))
elif t == 6:
    print(''.join([chr(97 + x % 26) for x in numbers]))
else:
    indexs = [False]*n
    idx = 0
    while True:
        if idx >= n:
            print('Out')
            break
        if indexs[idx]:
            print('Cyclic')
            break
        if idx == n-1:
            print('Done')
            break
        indexs[idx] = True
        idx = numbers[idx]
