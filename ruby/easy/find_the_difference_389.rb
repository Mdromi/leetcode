def find_the_difference(s, t)
    sum_s = s.chars.reduce(0) { |sum, char| sum + char.ord }
    sum_t = t.chars.reduce(0) { |sum, char| sum + char.ord }
    
    (sum_t - sum_s).chr
end

# Test cases
puts find_the_difference("abcd", "abcde") # Output: "e"
puts find_the_difference("", "y")         # Output: "y"
