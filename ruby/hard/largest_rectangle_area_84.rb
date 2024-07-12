def largest_rectangle_area(heights)
    stack = []
    max_area = 0
    index = 0
  
    while index < heights.length
      if stack.empty? || heights[index] >= heights[stack.last]
        stack.push(index)
        index += 1
      else
        top_of_stack = stack.pop
        area = heights[top_of_stack] * (stack.empty? ? index : index - stack.last - 1)
        max_area = [max_area, area].max
      end
    end
  
    while !stack.empty?
      top_of_stack = stack.pop
      area = heights[top_of_stack] * (stack.empty? ? index : index - stack.last - 1)
      max_area = [max_area, area].max
    end
  
    max_area
end

# Test cases
heights1 = [2, 1, 5, 6, 2, 3]
heights2 = [2, 4]

puts largest_rectangle_area(heights1) # Output: 10
puts largest_rectangle_area(heights2) # Output: 4
  