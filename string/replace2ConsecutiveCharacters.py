""" program to swaps first pair(1st and 2nd character) with second pair(3rd and 4th) for every quadruplet substring
And keep leftover string in same order
eg 1: input: ABCDEFGH
  output: CDABGHEF

  eg 2: input: ABCDEFGHI
        output: CDABGHEFI
"""

def iterative(s):
    result = ""
    while len(s) > 3:
        result += s[2]+s[3]+s[0]+s[1]
        s = s[4:]
    else:
        result += s
    return result

user_input=input("enter value: ")
output=iterative(user_input)
print(output)
