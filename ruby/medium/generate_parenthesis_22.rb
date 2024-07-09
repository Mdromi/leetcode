# @param {Integer} n
# @return {String[]}
def generate_parenthesis(n)
    result = []
    
    def backtrack(result, s = '', open = 0, close = 0, max)
      if s.length == max * 2
        result << s
        return
      end
      
      if open < max
        backtrack(result, s + '(', open + 1, close, max)
      end
      
      if close < open
        backtrack(result, s + ')', open, close + 1, max)
      end
    end
    
    backtrack(result, '', 0, 0, n)
    result
end

# Test cases
n1 = 3
n2 = 1

puts generate_parenthesis(n1).inspect # Output: ["((()))", "(()())", "(())()", "()(())", "()()()"]
puts generate_parenthesis(n2).inspect # Output: ["()"]
