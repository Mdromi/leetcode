def top_k_frequent(nums, k)
    # Step 1: Count frequencies
    freq_map = Hash.new(0)
    nums.each { |num| freq_map[num] += 1 }
  
    # Step 2: Create buckets
    max_freq = nums.size
    buckets = Array.new(max_freq + 1) { [] }
    freq_map.each { |num, freq| buckets[freq] << num }
  
    # Step 3: Collect top k elements
    result = []
    buckets.reverse_each do |bucket|
      bucket.each do |num|
        result << num
        return result if result.size == k
      end
    end
    result
  end
  