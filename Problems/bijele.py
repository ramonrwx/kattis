pieces = [1, 1, 2, 2, 2, 8]
missing_pieces = [0]*6

pieces_given = [int(n) for n in input().split()]

for i, (x, y) in enumerate(zip(pieces, pieces_given)):
    missing_pieces[i] = x-y

for p in missing_pieces:
    print(p, end=' ')
