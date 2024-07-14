def length_of_last_word(s)
    length = 0
    i = s.length - 1
    
    # Skip trailing spaces
    while i >= 0 && s[i] == ' '
      i -= 1
    end
    
    # Count the length of the last word
    while i >= 0 && s[i] != ' '
      length += 1
      i -= 1
    end
    
    length
end

# Test cases
puts length_of_last_word("Hello World")                # Output: 5
puts length_of_last_word("   fly me   to   the moon  ") # Output: 4
puts length_of_last_word("luffy is still joyboy")       # Output: 6
  