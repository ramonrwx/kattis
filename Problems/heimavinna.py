problems = input().split(';')

total = 0

for problem in problems:
    if problem.__contains__('-'):
        a, b = problem.split('-')
        total += int(b) - int(a)
    total += 1

print(total)
