def int_to_roman(num)
    # Roman numeral lookup table in descending order of value
    roman_map = [
      [1000, "M"], [900, "CM"], [500, "D"], [400, "CD"],
      [100, "C"], [90, "XC"], [50, "L"], [40, "XL"],
      [10, "X"], [9, "IX"], [5, "V"], [4, "IV"], [1, "I"]
    ]
  
    roman = ""
    roman_map.each do |value, symbol|
      while num >= value
        roman += symbol
        num -= value
      end
    end
  
    roman
  end
  
  # Test cases
  puts int_to_roman(3749)  # Output: "MMMDCCXLIX"
  puts int_to_roman(58)    # Output: "LVIII"
  puts int_to_roman(1994)  # Output: "MCMXCIV"
  