package main

import (
	"fmt"
	"sort"
)

// Approach 1: Using Sorting
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// Sort the string to create a key for the hashmap
		sortedStr := sortString(str)

		// Append the original string to the corresponding group in the hashmap
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	// Collect all groups of anagrams
	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// Helper function to sort a string and return its sorted version
func sortString(s string) string {
	sBytes := []byte(s)
	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})
	return string(sBytes)
}

func main() {
	// Test cases
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs1)) // Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	strs2 := []string{""}
	fmt.Println(groupAnagrams(strs2)) // Output: [[""]]

	strs3 := []string{"a"}
	fmt.Println(groupAnagrams(strs3)) // Output: [["a"]]
}
