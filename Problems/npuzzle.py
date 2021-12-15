board = [str, str, str, str]
locations = {
    'A': (0, 0), 'B': (0, 1),
    'C': (0, 2), 'D': (0, 3),
    'E': (1, 0), 'F': (1, 1),
    'G': (1, 2), 'H': (1, 3),
    'I': (2, 0), 'J': (2, 1),
    'K': (2, 2), 'L': (2, 3),
    'M': (3, 0), 'N': (3, 1),
    'O': (3, 2)
}

for i in range(4):
    board[i] = input()

total = 0

for i in range(4):
    for j in range(4):
        if board[i][j] != '.':
            total += abs(i - locations[board[i][j]][0]) + \
                abs(j - locations[board[i][j]][1])

print(total)
