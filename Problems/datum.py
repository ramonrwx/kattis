from datetime import date

days_of_week = [
    'Monday', 'Tuesday', 'Wednesday',
    'Thursday', 'Friday', 'Saturday',
    'Sunday'
]

day, month = (int(d) for d in input().split())
weekday = date(year=2009, month=month, day=day).weekday()

print(days_of_week[weekday])
