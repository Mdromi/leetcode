def trap(height)
    left = 0
    right = height.length - 1
    left_max = 0
    right_max = 0
    total_water = 0
  
    while left < right
      if height[left] < height[right]
        if height[left] >= left_max
          left_max = height[left]
        else
          total_water += left_max - height[left]
        end
        left += 1
      else
        if height[right] >= right_max
          right_max = height[right]
        else
          total_water += right_max - height[right]
        end
        right -= 1
      end
    end
  
    total_water
  end
  
# Test cases
puts trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) # Output: 6
puts trap([4, 2, 0, 3, 2, 5])                  # Output: 9
