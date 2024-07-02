# Method 1: Using String Manipulation
def is_palindrome(s)
    # Remove non-alphanumeric characters and convert to lowercase
    filtered_str = s.gsub(/[^0-9a-z]/i, '').downcase
  
    # Check if the filtered string is equal to its reverse
    filtered_str == filtered_str.reverse
  end
  
  # Test cases
  puts is_palindrome("A man, a plan, a canal: Panama")  # Output: true
  puts is_palindrome("race a car")                     # Output: false
  puts is_palindrome(" ")                              # Output: true
  
# Method 2: Using Two Pointers
def is_palindrome_two_pointer(s)
    left, right = 0, s.length - 1
  
    while left < right
      # Move left pointer to the next alphanumeric character
      left += 1 while left < right && !s[left].match?(/[0-9a-zA-Z]/)
      # Move right pointer to the previous alphanumeric character
      right -= 1 while left < right && !s[right].match?(/[0-9a-zA-Z]/)
      
      # Compare the characters
      if s[left].downcase != s[right].downcase
        return false
      end
  
      # Move pointers towards the center
      left += 1
      right -= 1
    end
  
    true
  end
  
  # Test cases
  puts is_palindrome_two_pointer("A man, a plan, a canal: Panama")  # Output: true
  puts is_palindrome_two_pointer("race a car")                     # Output: false
  puts is_palindrome_two_pointer(" ")                              # Output: true
  