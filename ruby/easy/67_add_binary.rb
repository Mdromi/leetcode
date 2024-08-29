def add_binary(a, b)
  # Convert binary strings to integers, add them, and convert the result back to binary
  (a.to_i(2) + b.to_i(2)).to_s(2)
end

# Example 1
a = "11"
b = "1"
puts add_binary(a, b) # Output: "100"

# Example 2
a = "1010"
b = "1011"
puts add_binary(a, b) # Output: "10101"
