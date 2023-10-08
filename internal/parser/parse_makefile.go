package internal

import (
	"bufio"
	"fmt"

	"io/fs"
	"strings"
)

type target struct {
	name         string
	dependencies []string
	command      string
}

func ParseMakeFile(fileSystem fs.FS, filename string) []target {
	file, _ := fileSystem.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targets := make([]target, 0)

	targetLine, commandLine := "", ""
	for scanner.Scan() {
		textline := scanner.Text()
		fmt.Println(len(textline))
		if len(textline)==0 {
			continue
		} else if strings.HasPrefix(textline, "    ") {
			commandLine = textline
			targets = append(targets, newTarget(targetLine, commandLine))
			targetLine, commandLine = "", ""
		} else {
			targetLine = textline
		}
	}
	return targets
}

func newTarget(targetLine string, commandLine string) target {
	target := target{}
	splittedTargetLine := strings.Split(targetLine, " ")
	target.name = splittedTargetLine[0]
	target.name= target.name[0:len(target.name)-1]
	target.dependencies = splittedTargetLine[1:]
	target.command = strings.TrimPrefix(commandLine,"    ")

	return target
}
