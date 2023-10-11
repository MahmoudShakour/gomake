package internal

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Target struct {
	Id           int
	Name         string
	Dependencies []string
	Command      string
}

func ParseMakeFile(fileSystem fs.FS, filename string) []Target {
	temp := 0
	file, _ := fileSystem.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targets := make([]Target, 0)

	targetLine, commandLine := "", ""
	for scanner.Scan() {
		textline := scanner.Text()
		if len(textline) == 0 {
			continue
		} else if strings.HasPrefix(textline, "    ") {
			commandLine = textline
			targets = append(targets, newTarget(targetLine, commandLine, temp))
			targetLine, commandLine = "", ""
			temp++
		} else {
			targetLine = textline
		}
	}
	return targets
}

func newTarget(targetLine string, commandLine string, id int) Target {
	target := Target{}
	splittedTargetLine := strings.Split(targetLine, " ")
	target.Name = splittedTargetLine[0]
	target.Name = target.Name[0 : len(target.Name)-1]
	target.Dependencies = splittedTargetLine[1:]
	target.Command = strings.TrimPrefix(commandLine, "    ")
	target.Id = id
	return target
}

func ParsePath(filePath string) (fs.FS, string) {
	directoryName := os.DirFS(filepath.Dir(filePath))
	fileName := filepath.Base(filePath)

	return directoryName, fileName
}
