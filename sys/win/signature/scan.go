package signature

import (
	"unsafe"
)

func Scan(start uintptr, length int64, p *Pattern) int64 {
	// get the length of the pattern
	patternLength := p.Length()
	// check if the pattern longer the the search
	if length < int64(patternLength) {
		return 0
	}
	// calculate the length of the scan
	scanLength := length - (int64(patternLength) - 1)

	// search for the pattern
	for i := int64(0); i < scanLength; i++ {
		// make the flag for found
		found := true
		for j := 0; j < patternLength; j++ {
			// calculate the ptr
			ptr := start + uintptr(i) + uintptr(j)
			// cast to bytePtr
			bytePtr := (*byte)(unsafe.Pointer(ptr))
			// match for the pattern
			flag := p.MatchAt(j, *bytePtr)

			// check if the flag matchs
			if !flag {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	// return not found with -1
	return -1
}
