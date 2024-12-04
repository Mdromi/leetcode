def is_palindrome(x)
    # Early return if the number is negative or divisible by 10 (except for 0)
    return false if x < 0 || (x % 10 == 0 && x != 0)
  
    reversed_half = 0
    while x > reversed_half
      reversed_half = reversed_half * 10 + x % 10
      x /= 10
    end
  
    # x is equal to reversed_half for even-length numbers, or x is equal to reversed_half / 10 for odd-length numbers
    x == reversed_half || x == reversed_half / 10
  end
  
  # Test cases
  puts is_palindrome(121)  # Output: true
  puts is_palindrome(-121) # Output: false
  puts is_palindrome(10)   # Output: false
  