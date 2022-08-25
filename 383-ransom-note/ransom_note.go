package lc

// Given two strings ransomNote and magazine,
// return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.

import "strings"

func CanConstruct(ransomNote string, magazine string) bool {

	for _, a := range ransomNote {
		// string.Count will count occurance of an character in a string
		// if the count matches then that means magazine contains the substring of ransomnote
		if strings.Count(ransomNote, string(a)) > strings.Count(magazine, string(a)) {
			return false
		}
	}
	return true
}
