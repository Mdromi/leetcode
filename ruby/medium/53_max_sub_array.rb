def max_sub_array1(nums)
    return 0 if nums.empty?
  
    current_sum = nums[0]
    max_sum = nums[0]
  
    nums[1..].each do |num|
      current_sum = [num, current_sum + num].max
      max_sum = [max_sum, current_sum].max
    end
  
    max_sum
end
  

def max_product(nums)
    # Initialize the variables
    max_product = nums[0]
    min_product = nums[0]
    result = nums[0]
  
    # Iterate through the array starting from the second element
    nums[1..].each do |current|
      # If the current number is negative, swap max and min
      if current < 0
        max_product, min_product = min_product, max_product
      end
  
      # Update max_product and min_product
      max_product = [current, max_product * current].max
      min_product = [current, min_product * current].min
  
      # Update the result to track the maximum product seen so far
      result = [result, max_product].max
    end
  
    result
  end
  