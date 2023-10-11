package internal

import (
	"os/exec"
	"strings"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

func BuildTopoSort(node int, visited []int, adjList map[int][]int) []int {
	visited[node] = 1
	topoSort := make([]int, 0)
	for _, to := range adjList[node] {
		if visited[to] == 0 {
			topoSort = append(topoSort, BuildTopoSort(to, visited, adjList)...)
		}
	}
	topoSort = append(topoSort, node)
	return topoSort
}

func ExecuteTarget(targets []internal.Target, targetId int, numberOfTargets int) (string, error) {
	err := CheckInvalidDependency(targets)
	if err != nil {
		return "", err
	}
	err = CheckCyclicDependency(targets)
	if err != nil {
		return "", err
	}
	targetIdToTargetCommand := map[int]string{}
	for _, target := range targets {
		targetIdToTargetCommand[target.Id] = target.Command
	}
	
	adjList := BuildAdjList(targets)
	visited := make([]int, numberOfTargets)
	topoSort := BuildTopoSort(targetId, visited, adjList)
	
	output := ""
	for _, targetId := range topoSort {
		targetCommand := targetIdToTargetCommand[targetId]
		commandOutput, err := executeCommand(targetCommand)
		if err != nil {
			return "", err
		} else if len(commandOutput) != 0 {
			output += commandOutput
		}
	}
	return output, nil
}

func executeCommand(command string) (string, error) {
	splittedCommand := strings.Split(command, " ")

	cmd := exec.Command(splittedCommand[0], splittedCommand[1:]...)
	output, err := cmd.CombinedOutput()

	return string(output), err
}
