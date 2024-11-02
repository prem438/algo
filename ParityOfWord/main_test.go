package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestParityOfWord_1(t *testing.T) {
	nums := [][2]uint64{
		{0b1101, 1},
		{0b1111, 0},
		{0b011111, 1},
	}
	for i, n := range nums {
		logger.Println("--------test ", i+1)

		r := ParityOfWord(n[0])
		assert.Equal(t, r, uint8(n[1]))
		logger.Printf("num 0b%b, expected %d result %d\n", n[0],n[1], r)

	}
}
