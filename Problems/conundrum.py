text = input()

TARGET = 'PER'
count_days = 0

for i, char in enumerate(text):
    if char != TARGET[i % 3]:
        count_days += 1
print(count_days)
