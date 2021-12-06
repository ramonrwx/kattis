from re import finditer

smiles = input()

smile_regex = r':\)|;\)|:-\)|;-\)'
for match in finditer(smile_regex, smiles):
    print(match.start())
