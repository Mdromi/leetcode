def is_happy(n)
    seen = {}
    
    while n != 1 && !seen[n]
      seen[n] = true
      n = sum_of_squares(n)
    end
    
    n == 1
end
  
def sum_of_squares(n)
  sum = 0
  while n > 0
    digit = n % 10
    sum += digit * digit
    n /= 10
  end
  sum
end

# Test cases
n1 = 19
n2 = 2

puts is_happy(n1) # Output: true
puts is_happy(n2) # Output: false
  