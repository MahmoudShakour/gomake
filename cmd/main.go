package main

import (
	"flag"
	"fmt"
	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

func getArguments() (string,string) {
	makeFilePath:=flag.String("f","./makefile","path of the makefile")
	target:=flag.String("t","all","target")
	flag.Parse()
	return *makeFilePath,*target
}


func main() {
	
	
	makeFilePath,_:=getArguments()
	directoryName,fileName := internal.ParsePath(makeFilePath)
	targets:=internal.ParseMakeFile(directoryName,fileName)

	
	fmt.Println(targets)
	
}