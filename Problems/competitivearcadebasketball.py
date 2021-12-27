n, p, m = (int(x) for x in input().split())
players = {input(): 0 for _ in range(n)}

for _ in range(m):
    name, points = input().split()
    if name in players:
        players[name] += int(points)

winners = []

for player, points in players.items():
    if points >= p:
        print(f'{player} wins!')
        winners.append(winners)

if len(winners) <= 0:
    print('No winner!')
