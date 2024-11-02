package main

import "fmt"

func Check(t any){
	_, ok := t.(bool)
	fmt.Println(ok)
}


func  main()  {
	Check(true)
	
}