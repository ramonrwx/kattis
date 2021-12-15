cards = {}

for card in input().split():
    cards.setdefault(card[0], 0)
    cards[card[0]] += 1
print(max(cards.values()))
