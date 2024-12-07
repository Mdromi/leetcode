def product_except_self(nums)
    n = nums.length
    answer = Array.new(n, 1)
  
    # Step 1: Compute prefix products
    prefix = 1
    (0...n).each do |i|
      answer[i] = prefix
      prefix *= nums[i]
    end
  
    # Step 2: Compute suffix products and multiply
    suffix = 1
    (n - 1).downto(0) do |i|
      answer[i] *= suffix
      suffix *= nums[i]
    end
  
    answer
  end
  