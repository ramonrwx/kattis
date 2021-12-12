spam = input()
whitespaces, lowercases, uppercases, symbols = [0, 0, 0, 0]

for char in spam:
    if char == '_':
        whitespaces += 1
    elif char.islower():
        lowercases += 1
    elif char.isupper():
        uppercases += 1
    else:
        symbols += 1

print(whitespaces/len(spam))
print(lowercases/len(spam))
print(uppercases/len(spam))
print(symbols/len(spam))
