# Example 1: Basic Calculator
def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

def divide(a, b):
    if b != 0:
        return a / b
    else:
        print("Error: Division by zero is not allowed!")
        return None

# Example 2: Factorial Function
def factorial(n):
    if n == 0 or n == 1:
        return 1
    else:
        return n * factorial(n-1)

# Example 3: While Loop
n = 5
while n > 0:
    print(n)
    n -= 1

# Example 4: For Loop with List of Numbers
numbers = [2, 4, 6, 8]
for number in numbers:
    if number % 2 == 0:
        print(number)

# Example 5: Function to Sort an Array
def bubble_sort(arr):
    n = len(arr)
    for i in range(n-1):
        for j in range(n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
