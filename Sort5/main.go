package main

import (
	"log"
	"sort"
)

var logger *log.Logger

// sortN algorithm
func sortN(A []int, N int) []int {
	if len(A) > N {
		logger.Fatal("length is more than 5")
	}
	sort.Ints(A)
	logger.Print(A)
	return A
}

func init(){
    logger = log.Default()
	logger.SetFlags(log.Lshortfile)
}

func getMaxN(A []int, N, C int) {

	for len(A) > C {
		P := []([]int){}
		logger.Println("---loop---")

		for i := 0; i < len(A); i = i + N {
			t := A[i:]

			if len(t) > N {
				P = append(P, sortN(t[:N], N))
			} else {
				P = append(P, sortN(t, N))
			}
		}

		A = []int{}

		for _, p := range P {
			n := C
			if len(p) < C {
				n = len(p)
			}
			A = append(A, p[:n]...)
		}
	}
	logger.Print("result", A)
}