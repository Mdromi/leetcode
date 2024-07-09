def daily_temperatures(temperatures)
    n = temperatures.length
    answer = Array.new(n, 0)
    stack = []
  
    (0...n).each do |i|
      while !stack.empty? && temperatures[i] > temperatures[stack.last]
        idx = stack.pop
        puts "idx - #{idx}"
        puts "i - #{i}"
        answer[idx] = i - idx
      end
      stack.push(i)
    end
  
    answer
end

# Test cases
temperatures1 = [73, 74, 75, 71, 69, 72, 76, 73]
temperatures2 = [30, 40, 50, 60]
temperatures3 = [30, 60, 90]

puts daily_temperatures(temperatures1).inspect  # Output: [1, 1, 4, 2, 1, 1, 0, 0]
puts daily_temperatures(temperatures2).inspect  # Output: [1, 1, 1, 0]
puts daily_temperatures(temperatures3).inspect  # Output: [1, 1, 0]
  