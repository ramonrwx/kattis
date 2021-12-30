cia_blimps = []

for i in range(5):
    blimp = input()
    if 'FBI' in blimp:
        cia_blimps.append(str(i + 1))

if cia_blimps:
    print(' '.join(cia_blimps))
else:
    print('HE GOT AWAY!')
