def is_power_of_two(n)
  return false if n <= 0
  
  while n % 2 == 0
    n /= 2
  end
  
  n == 1
end

puts is_power_of_two(1)  # Output: true
puts is_power_of_two(16) # Output: true
puts is_power_of_two(3)  # Output: false
