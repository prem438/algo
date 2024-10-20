package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func init(){
	fmt.Println("i am test init")
}

func TestSort5(t *testing.T) {
	t.Log("running test")
	A := []int{}
	N := 5
	C := 3

	for i := N; i > 0; i-- {
		A = append(A, i)
	}

	getMaxN(A,N,C)
	assert.Equal(t, nil,nil)
}