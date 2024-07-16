def group_anagrams(strs)
    anagrams = Hash.new { |hash, key| hash[key] = [] }
  
    strs.each do |str|
      sorted_str = str.chars.sort.join
      anagrams[sorted_str] << str
    end
  
    anagrams.values
end

# Test cases
p group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]) # Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
p group_anagrams([""])                                       # Output: [[""]]
p group_anagrams(["a"])                                      # Output: [["a"]]
  