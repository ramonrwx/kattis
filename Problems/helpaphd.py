cases = int(input())
problems = [input() for _ in range(cases)]

for i in range(cases):
    if problems[i] != 'P=NP':
        a, b = (int(x) for x in problems[i].split('+'))
        print(a + b)
    else:
        print('skipped')
    # if problems[i] != 'P=NP':
    #    print(eval(problems[i]))
    # else:
    #    print('skipped')
