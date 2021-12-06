cases = int(input())
words = [input() for i in range(cases)]

for word in words:
    output = ''

    for i in range(len(word)):
        sqrt = (len(word) ** (1/2))

        index = int((i % sqrt+1)*sqrt - i // sqrt - 1)
        output += word[index]

    print(output)
