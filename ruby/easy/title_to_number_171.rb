def title_to_number(columnTitle)
    result = 0
    columnTitle.each_char do |char|
      result = result * 26 + (char.ord - 'A'.ord + 1)
    end
    result
end

# Test cases
puts title_to_number("A")  # Output: 1
puts title_to_number("AB") # Output: 28
puts title_to_number("ZY") # Output: 701
