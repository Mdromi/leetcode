def is_valid(s)
    stack = []
    pairs = { ')' => '(', '}' => '{', ']' => '[' }
  
    s.each_char do |char|
      if pairs.values.include?(char)  # If it's an opening bracket
        stack.push(char)
      elsif pairs.keys.include?(char)  # If it's a closing bracket
        return false if stack.empty? || stack.pop != pairs[char]
      end
    end
  
    stack.empty?
end
  
# Test cases
puts is_valid("()")        # Output: true
puts is_valid("()[]{}")    # Output: true
puts is_valid("(]")        # Output: false
puts is_valid("([)]")      # Output: false
puts is_valid("{[]}")      # Output: true
