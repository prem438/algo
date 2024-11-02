package main

import "log"

var logger *log.Logger
func init(){
	logger = log.Default()
	logger.SetFlags(log.Lshortfile)
}

func ParityOfWord(x uint64) uint8{
    s := 64
	for i := 1 ; x > 1; i++ {
		s = s/2 
		x ^= (x >> s)

		//create 8 bit mask
		mask := uint64(1<<s)-1
		x &= mask

		
		logger.Printf("0b%b ,shift %d, m 0x%x\n", x, s, mask)
	}
	return uint8(x)
}