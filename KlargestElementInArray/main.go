package main

import (
	"log"
)

var logger = log.Default()

func init() {
	logger.SetFlags(log.Lshortfile)

}

func findKthMaxElement() {
	logger.Print("hello")
}
