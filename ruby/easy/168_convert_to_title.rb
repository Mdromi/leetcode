def convert_to_title(column_number)
  result = ""
  
  while column_number > 0
    # Adjust for 0-based index
    column_number -= 1
    # Compute the letter for the current digit
    result = (65 + column_number % 26).chr + result
    # Move to the next digit
    column_number /= 26
  end
  
  result
end

# Examples
puts convert_to_title(1)    # Output: "A"
puts convert_to_title(28)   # Output: "AB"
puts convert_to_title(701)  # Output: "ZY"
