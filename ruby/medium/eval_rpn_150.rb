def eval_rpn(tokens)
    stack = []
  
    tokens.each do |token|
      if ['+', '-', '*', '/'].include?(token)
        b = stack.pop
        a = stack.pop
        case token
        when '+'
          stack.push(a + b)
        when '-'
          stack.push(a - b)
        when '*'
          stack.push(a * b)
        when '/'
          stack.push((a.to_f / b).to_i)  # Ensure truncation towards zero
        end
      else
        stack.push(token.to_i)
      end
    end
  
    stack.pop
end

# Test cases
tokens1 = ["2", "1", "+", "3", "*"]
tokens2 = ["4", "13", "5", "/", "+"]
tokens3 = ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]

puts eval_rpn(tokens1) # Output: 9
puts eval_rpn(tokens2) # Output: 6
puts eval_rpn(tokens3) # Output: 22
  