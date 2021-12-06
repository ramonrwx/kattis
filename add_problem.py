#!/usr/bin/env python
from shutil import copy

KATTIS_URL = 'https://open.kattis.com/problems/'

success = True

prob_title = input('Enter the problem title: ')
prob_url = input('Enter the problem URL: ')
file_name, ext = input('Enter the filename: ').split('.')

if ext == 'go':
    lang = 'Go'
elif ext == 'py':
    lang = 'Python'

if prob_title and prob_url:
    file = open('./README.md', 'a')
    file.write(
        f'| [{prob_title}]({KATTIS_URL}{prob_url}) | [{lang}](Problems/{file_name}.{ext}) |\n')
else:
    success = False

if success:
    copy(f'./main.{ext}', f'./Problems/{file_name}.{ext}',
         follow_symlinks=True)
    print(f'The {prob_title} problem was added to the README.md')
else:
    print('Okay, Houston, we\'ve had a problem here')
