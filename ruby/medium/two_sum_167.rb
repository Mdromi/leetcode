def two_sum(numbers, target)
    left = 0
    right = numbers.length - 1
  
    while left < right
      sum = numbers[left] + numbers[right]
      if sum == target
        return [left + 1, right + 1]
      elsif sum < target
        left += 1
      else
        right -= 1
      end
    end
  end
  
# Test cases
puts two_sum([2, 7, 11, 15], 9).inspect  # Output: [1, 2]
puts two_sum([2, 3, 4], 6).inspect       # Output: [1, 3]
puts two_sum([-1, 0], -1).inspect        # Output: [1, 2]
  