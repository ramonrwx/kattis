hour, minutes = [int(x) for x in input().split()]
minutes -= 45

if minutes < 0:
    hour -= 1
    minutes += 60
if hour < 0:
    hour = 23

print(hour, minutes)
