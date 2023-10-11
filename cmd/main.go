package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	exec "github.com/MahmoudShakour/gomake.git/internal/TargetUtility"
	parser "github.com/MahmoudShakour/gomake.git/internal/parser"
)

var (
	defaultFilePath   = "/home/mahmoudshakour/Workspace/gomake/makefile"
	defaultTarget     = "run"
	ErrTargetNotFound = errors.New("given target is not found")
)

func getArguments() (string, string) {
	makeFilePath := flag.String("f", defaultFilePath, "")
	target := flag.String("t", defaultTarget, "")
	flag.Parse()
	return *makeFilePath, *target
}

func getTargetId(targets []parser.Target, targetName string) (int, error) {
	for _, target := range targets {
		if target.Name == targetName {
			return target.Id, nil
		}
	}
	return -1, ErrTargetNotFound
}
func main() {

	makeFilePath, target := getArguments()
	directoryName, fileName := parser.ParsePath(makeFilePath)
	targets := parser.ParseMakeFile(directoryName, fileName)
	targetId, err := getTargetId(targets, target)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	output, err := exec.ExecuteTarget(targets, targetId, len(targets))

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		fmt.Print(output)
	}
}
