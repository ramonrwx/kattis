#!/usr/bin/env python
from pathlib import Path

KATTIS_URL = 'https://open.kattis.com/problems/'
ROOT = Path(__file__).resolve().parents[1]

success = True

prob_title = input('Enter the problem title: ')
prob_url = input('Enter the problem URL: ')
same_name = input('The url of the problem is the same name as the file? [Y/n] ')

if same_name == '' or same_name == 'y':
    problem = prob_url
else:
    file_name = input('Enter the file name: ')
    problem = file_name

if prob_title and problem:
    file = open(ROOT.joinpath('./README.md'), 'a')
    file.write(f'| [{prob_title}]({KATTIS_URL}{problem}) | [Go](Go/{problem}.go) |\n')
else:
    success = False

if success:
    print(f'The {prob_title} problem was added to the README.md')
else:
    print('Okay, Houston, we\'ve had a problem here')
