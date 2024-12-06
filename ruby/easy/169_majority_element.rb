def majority_element(nums)
    candidate = nil
    count = 0
  
    # Phase 1: Find a candidate
    nums.each do |num|
      if count == 0
        candidate = num
      end
      count += (num == candidate ? 1 : -1)
    end
  
    # Since the problem guarantees that a majority element always exists,
    # we can directly return the candidate.
    candidate
end
  