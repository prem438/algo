package main

func findNumbersOfSetBits(x uint64) uint64 {
	r := uint64(0)
	for x > 0 {
		r++
		x &= (x - 1)
	}
	return r
}
