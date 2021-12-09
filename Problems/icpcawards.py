cases = int(input())
institutions = []

for _ in range(cases):
    institution, team = input().split()
    if len(institutions) >= 12:
        break
    elif institution not in institutions:
        print(institution, team)
        institutions.append(institution)
