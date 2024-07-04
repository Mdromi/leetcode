def max_area(height)
    left = 0
    right = height.length - 1
    max_area = 0
  
    while left < right
      width = right - left
      current_height = [height[left], height[right]].min
      current_area = width * current_height
      max_area = [max_area, current_area].max
  
      if height[left] < height[right]
        left += 1
      else
        right -= 1
      end
    end
  
    max_area
end
  
# Test cases
puts max_area([1, 8, 6, 2, 5, 4, 8, 3, 7])  # Output: 49
puts max_area([1, 1])                      # Output: 1
puts max_area([4, 3, 2, 1, 4])             # Output: 16
puts max_area([1, 2, 1])                   # Output: 2
