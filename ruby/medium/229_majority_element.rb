def majority_element(nums)
    return nums if nums.length <= 1
  
    # Step 1: Identify up to two potential candidates
    candidate1, candidate2 = nil, nil
    count1, count2 = 0, 0
  
    nums.each do |num|
      if num == candidate1
        count1 += 1
      elsif num == candidate2
        count2 += 1
      elsif count1 == 0
        candidate1 = num
        count1 = 1
      elsif count2 == 0
        candidate2 = num
        count2 = 1
      else
        count1 -= 1
        count2 -= 1
      end
    end
  
    # Step 2: Verify the candidates
    result = []
    n = nums.length
    count1 = nums.count { |num| num == candidate1 }
    count2 = nums.count { |num| num == candidate2 }
  
    result << candidate1 if count1 > n / 3
    result << candidate2 if count2 > n / 3
  
    result
  end
  