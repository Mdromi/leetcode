# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
      @val = val
      @left = left
      @right = right
    end
end

# Function to convert sorted array to BST
def sorted_array_to_bst(nums)
  return nil if nums.empty?

  construct_bst(nums, 0, nums.length - 1)
end

def construct_bst(nums, left, right)
  return nil if left > right

  mid = (left + right) / 2
  root = TreeNode.new(nums[mid])
  root.left = construct_bst(nums, left, mid - 1)
  root.right = construct_bst(nums, mid + 1, right)
  
  root
end

# Helper function to print the tree (level order)
def print_tree(node)
  return "[]" if node.nil?

  result = []
  queue = [node]
  while !queue.empty?
    current = queue.shift
    if current
      result << current.val
      queue << current.left
      queue << current.right
    else
      result << "null"
    end
  end

  # Remove trailing nulls
  while result.last == "null"
    result.pop
  end

  result.inspect
end

# Test cases
nums1 = [-10, -3, 0, 5, 9]
nums2 = [1, 3]

root1 = sorted_array_to_bst(nums1)
root2 = sorted_array_to_bst(nums2)

puts print_tree(root1) # Output: [0, -3, 9, -10, null, 5]
puts print_tree(root2) # Output: [3, 1]
  