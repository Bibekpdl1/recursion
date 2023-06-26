# program to take user input and replace pi with 3.14
def convert_pi(x):
  return x.replace("pi","3.14")

user_input = input("Enter your word: ")
print("user input :",user_input)
output= convert_pi(user_input)
print("After conversion: ",output)
