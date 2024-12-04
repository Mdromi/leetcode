def remove_element(nums, val)
    write = 0 # Pointer for placing non-val elements
  
    nums.each do |num|
      if num != val
        nums[write] = num # Write the valid element
        write += 1
      end
    end
  
    write # The number of valid elements
  end
  
  # Examples
  nums1 = [3, 2, 2, 3]
  k1 = remove_element(nums1, 3)
  puts "Output: #{k1}, nums: #{nums1[0...k1].inspect}" # Output: 2, nums: [2, 2]
  
  nums2 = [0, 1, 2, 2, 3, 0, 4, 2]
  k2 = remove_element(nums2, 2)
  puts "Output: #{k2}, nums: #{nums2[0...k2].inspect}" # Output: 5, nums: [0, 1, 3, 0, 4]
  