# Approach 1: Sorting
def is_anagram1(s, t)
    s.chars.sort == t.chars.sort
end

# Approach 2: Frequency Count
def is_anagram(s, t)
    return false unless s.length == t.length
  
    count = Hash.new(0)
  
    s.each_char { |char| count[char] += 1 }
    t.each_char { |char| count[char] -= 1 }
  
    count.values.all? { |val| val == 0 }
end

# Test cases
puts is_anagram("anagram", "nagaram") # Output: true
puts is_anagram("rat", "car")         # Output: false
