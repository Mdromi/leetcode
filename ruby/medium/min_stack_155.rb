class MinStack
    def initialize
      @stack = []
      @min_stack = []
    end
  
    def push(val)
      @stack.push(val)
      if @min_stack.empty? || val <= @min_stack[-1]
        @min_stack.push(val)
      end
    end
  
    def pop
      if @stack.pop == @min_stack[-1]
        @min_stack.pop
      end
    end
  
    def top
      @stack[-1]
    end
  
    def get_min
      @min_stack[-1]
    end
end
  
# Example usage
min_stack = MinStack.new
min_stack.push(-2)
min_stack.push(0)
min_stack.push(-3)
puts min_stack.get_min  # Output: -3
min_stack.pop
puts min_stack.top      # Output: 0
puts min_stack.get_min  # Output: -2
