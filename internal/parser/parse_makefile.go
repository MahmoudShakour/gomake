package internal

import (
	"bufio"
	"os"
	"path/filepath"
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


func ParsePath(filePath string) (fs.FS,string){
	directoryName:=os.DirFS(filepath.Dir(filePath))
	fileName:=filepath.Base(filePath)

	return directoryName,fileName
}