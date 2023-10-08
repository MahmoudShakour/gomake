package main

import (
	"flag"
	"fmt"
	
)


func main() {
	
	makeFilePath:=flag.String("f","./makefile","path of the makefile")
	target:=flag.String("t","all","target")
	flag.Parse()
	

	fmt.Println(*makeFilePath,*target)
	
}