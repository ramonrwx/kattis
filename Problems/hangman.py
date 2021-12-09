word_to_be_guessed = set(input())
guesses_given = input()
count_incorrect = 0

for character in guesses_given:
    if not word_to_be_guessed:
        print('WIN')
        break
    elif count_incorrect == 10:
        print('LOSE')
        break
    elif character in word_to_be_guessed:
        word_to_be_guessed.remove(character)
    else:
        count_incorrect += 1
