def two_sum(nums, target)
    hash = {} # Initialize a hash to store numbers and their indices
  
    nums.each_with_index do |num, index|
      complement = target - num
      # Check if the complement exists in the hash map
      if hash.key?(complement)
        return [hash[complement], index] # Return indices of the complement and current number
      end
      # Store the current number and its index in the hash map
      hash[num] = index
    end
  
    [] # If no solution exists (won't happen as per problem statement)
  end
  
  # Examples
  puts two_sum([2, 7, 11, 15], 9).inspect  # Output: [0, 1]
  puts two_sum([3, 2, 4], 6).inspect       # Output: [1, 2]
  puts two_sum([3, 3], 6).inspect          # Output: [0, 1]
  