gold, silver, copper = (int(x) for x in input().split())

power = 3 * gold + 2 * silver + 1 * copper

if power >= 8:
    print('Province or', end=' ')
elif power >= 5:
    print('Duchy or', end=' ')
elif power >= 2:
    print('Estate or', end=' ')

if power >= 6:
    print('Gold')
elif power >= 3:
    print('Silver')
else:
    print('Copper')
