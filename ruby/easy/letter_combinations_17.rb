def letter_combinations(digits)
    # Return an empty array if the input digits string is empty
    return [] if digits.empty?
  
    # Mapping of digits to corresponding letters
    digit_to_letters = {
      '2' => 'abc', '3' => 'def', '4' => 'ghi', '5' => 'jkl', '6' => 'mno',
      '7' => 'pqrs', '8' => 'tuv', '9' => 'wxyz'
    }
  
    result = []
  
    # Helper method to perform backtracking
    def backtrack(combination, next_digits, digit_to_letters, result)
      # If there are no more digits to check, the combination is complete
      if next_digits.empty?
        result << combination
      else
        # Iterate over all letters which map the next available digit
        digit_to_letters[next_digits[0]].each_char do |letter|
          # Append the current letter to the combination and proceed to the next digits
          backtrack(combination + letter, next_digits[1..], digit_to_letters, result)
        end
      end
    end
  
    # Initiate the backtracking process
    backtrack("", digits, digit_to_letters, result)
    result
end

# Test cases
puts letter_combinations("23").inspect  # Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
puts letter_combinations("").inspect    # Output: []
puts letter_combinations("2").inspect   # Output: ["a","b","c"]
  