package lc

// Given a roman numeral, convert it to an integer.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

func RomanToInt(s string) int {

	romans := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sLen := len(s)
	// assign last element as the result and use later
	response := romans[s[sLen-1]]
	//loop through the string from last but one element i.e. in reverse order
	for i := sLen - 2; i >= 0; i-- {
		// check if the current element is less than next element of string in the romans map
		// if current elemet is greater then sum it from response else do the difference
		if romans[s[i]] < romans[s[i+1]] {
			response -= romans[s[i]]
		} else {
			response += romans[s[i]]
		}
	}
	return response
}
