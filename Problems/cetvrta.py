col_x, xs = [], set()
col_y, ys = [], set()

for i in range(3):
    x, y = (int(n) for n in input().split())
    col_x.append(x)
    col_y.append(y)

    if col_x[i] not in xs:
        xs.add(col_x[i])
    else:
        xs.remove(col_x[i])

    if col_y[i] not in ys:
        ys.add(col_y[i])
    else:
        ys.remove(col_y[i])

print(xs.pop(), ys.pop())
