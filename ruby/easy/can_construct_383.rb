def can_construct(ransom_note, magazine)
    # Create a hash to count the frequency of each character in the magazine
    magazine_count = Hash.new(0)
    
    # Count each character in the magazine
    magazine.each_char do |char|
      magazine_count[char] += 1
    end
    
    # Check if each character in the ransom note can be constructed from the magazine
    ransom_note.each_char do |char|
      return false if magazine_count[char] <= 0
      magazine_count[char] -= 1
    end
    
    return true
end

# Test cases
puts can_construct("a", "b")         # Output: false
puts can_construct("aa", "ab")       # Output: false
puts can_construct("aa", "aab")      # Output: true
  