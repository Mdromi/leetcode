package main

import (
	"fmt"
)

// Approach 2: Using Character Frequency Counting
func groupAnagramsCharFreq(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		freq := countFrequency(str)
		key := getKey(freq)
		anagrams[key] = append(anagrams[key], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func countFrequency(s string) [26]int {
	freq := [26]int{}
	for _, ch := range s {
		freq[ch-'a']++
	}
	return freq
}

func getKey(freq [26]int) string {
	key := ""
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			key += fmt.Sprintf("%c%d", 'a'+i, freq[i])
		}
	}
	return key
}

func main() {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsCharFreq(strs1)) // Output: [["eat","tea","ate"],["tan","nat"],["bat"]]

	strs2 := []string{""}
	fmt.Println(groupAnagramsCharFreq(strs2)) // Output: [[""]]

	strs3 := []string{"a"}
	fmt.Println(groupAnagramsCharFreq(strs3)) // Output: [["a"]]
}
