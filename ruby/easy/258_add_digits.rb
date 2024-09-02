def add_digits(num)
  return num == 0 ? 0 : 1 + (num - 1) % 9
end

# Test cases
puts add_digits(38) # Output: 2
puts add_digits(0)  # Output: 0
