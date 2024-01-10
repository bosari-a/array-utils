package search

import "sort"

// Struct for a int array with binary search method and extra properties
type binarySearchInt struct {
	intArray []int
	low      int
	high     int
	mid      int
}

// Initialize a struct with binary search method for int array
// The struct contains an array that will be sorted based on the sorted boolean param
func BinaryIntFactory(a []int, sorted bool) *binarySearchInt {
	var bi binarySearchInt
	if !sorted {
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
	}

	bi.intArray = a
	l := len(a)
	bi.high = l - 1
	bi.low = 0
	bi.mid = (l - 1) / 2
	return &bi
}

// Method to search int array using binary search
// Returns the index of the element in the array if it is found
// Returns -1 if the element is not found
func (a *binarySearchInt) BinarySearchInt(find int) int {
	p := a.intArray[a.mid]
	if find == p {
		return a.mid
	} else if a.low == a.mid && a.high == a.mid {
		return -1
	}
	if find > p {
		a.low = a.mid + 1
		a.mid = (a.low + a.high) / 2
		return a.BinarySearchInt(find)
	}
	a.high = a.mid - 1
	a.mid = (a.low + a.high) / 2
	return a.BinarySearchInt(find)
}

// Struct for a string array with binary search method and extra properties
type binarySearchStr struct {
	strArray []string
	low      int
	high     int
	mid      int
}

// Initialize a struct with binary search method for string array
// The struct contains an array that will be sorted based on the sorted boolean param
func BinaryStrFactory(s []string, sorted bool) *binarySearchStr {
	var bs binarySearchStr
	if !sorted {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}

	bs.strArray = s
	l := len(s)
	bs.high = l - 1
	bs.low = 0
	bs.mid = (l - 1) / 2
	return &bs
}

// Method that implements binary search on an array of strings.
// The modifyStr function will modify the string being searched against (not the find string)
// so that a comparison happens between a segment of the search strings and the string you want to
// find. Can be useful when parsing lines without needing to create a new equally sized
// array of search strings.
func (bs *binarySearchStr) BinarySearchStr(find string, modifyStr func(string) string) int {
	s := bs.strArray[bs.mid]
	s = modifyStr(s)
	if find == s {
		return bs.mid
	} else if bs.low >= bs.high {
		return -1
	}
	if find > s {
		bs.low = bs.mid + 1
		bs.mid = (bs.low + bs.high) / 2
		return bs.BinarySearchStr(find, modifyStr)
	}
	bs.high = bs.mid - 1
	bs.mid = (bs.low + bs.high) / 2
	return bs.BinarySearchStr(find, modifyStr)
}
