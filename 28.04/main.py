def product_of_digits(n):
    """Returns the product of the digits of n."""
    result = 1
    for digit in str(n):
        result *= int(digit)
    return result

def find_largest_pandigital_number():
    """Finds the largest pandigital number with the property described in the problem."""
    # Initialize a dictionary to store the product of digits for each number
    products = {}
    # Loop through numbers in reverse order
    for i in range(987654321, 11, -1):
        # Check if the number has already been processed
        if i in products:
            tmp = products[i]
        else:
            # Compute the product of digits for the number
            tmp = product_of_digits(i)
            products[i] = tmp
        # If the product is zero, skip to the next number
        if tmp == 0:
            continue
        # Keep computing the product of digits until a single digit is reached
        while tmp > 9:
            tmp = product_of_digits(tmp)
            products[i] = tmp
            # If the product is zero, skip to the next number
            if tmp == 0:
                break
        # If the product is zero, skip to the next number
        if tmp == 0:
            continue
        # Check if the number is pandigital
        if sorted(str(i)) == [str(digit) for digit in range(1, len(str(i))+1)]:
            return i
    return 0

# Call the function to find the largest pandigital number
result = find_largest_pandigital_number()
print(result)
