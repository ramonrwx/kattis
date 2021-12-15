from string import ascii_uppercase
from operator import add, sub

encrypted_msg = input()
key = input()

ALPHABET = ascii_uppercase
decrypted_msg = ''

for i in range(len(encrypted_msg)):
    opr = sub if i % 2 == 0 else add
    decrypted_msg += ALPHABET[opr(ALPHABET
                                  .index(encrypted_msg[i]),
                                  ALPHABET.index(key[i])) % 26]
print(decrypted_msg)
