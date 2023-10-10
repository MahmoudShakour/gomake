package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"io/fs"
	"strings"
)

type Target struct {
	Name         string
	Dependencies []string
	Command      string
}

func ParseMakeFile(fileSystem fs.FS, filename string) []Target {
	file, _ := fileSystem.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targets := make([]Target, 0)

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

func newTarget(targetLine string, commandLine string) Target {
	target := Target{}
	splittedTargetLine := strings.Split(targetLine, " ")
	target.Name = splittedTargetLine[0]
	target.Name= target.Name[0:len(target.Name)-1]
	target.Dependencies = splittedTargetLine[1:]
	target.Command = strings.TrimPrefix(commandLine,"    ")

	return target
}


func ParsePath(filePath string) (fs.FS,string){
	directoryName:=os.DirFS(filepath.Dir(filePath))
	fileName:=filepath.Base(filePath)

	return directoryName,fileName
}