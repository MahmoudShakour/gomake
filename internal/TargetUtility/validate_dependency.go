package internal

import (
	"errors"
	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

var ErrDependencyNotFound = errors.New("given dependency is not found")
var ErrCyclicDependencyFound = errors.New("cyclic dependency is found")

func CheckInvalidDependency(targets []internal.Target) error {
	isTarget := map[string]int{}

	for _, target := range targets {
		isTarget[target.Name] = 1
	}

	for _, target := range targets {
		for _, dependency := range target.Dependencies {
			if isTarget[dependency] == 0 {
				return ErrDependencyNotFound
			}
		}
	}
	return nil
}

func CheckCyclicDependency(targets []internal.Target) error {
	numberOfTargets := len(targets)
	adjList:=BuildAdjList(targets)
	
	visited := make([]int, numberOfTargets)

	for i := 0; i < numberOfTargets; i++ {
		if visited[i] == 0 {
			cycleFound := dfs(i, adjList, visited)
			
			if cycleFound == true {
				return ErrCyclicDependencyFound
			}
		}
	}

	return nil
}

func dfs(node int, adjList map[int][]int, visited []int) bool {
	if visited[node] == 2 {
		return false
	}
	if visited[node] == 1 {
		return true
	}
	visited[node] = 1
	cycleFound := false
	for _, to := range adjList[node] {
		cycleFound = cycleFound || dfs(to, adjList, visited)
	}
	visited[node] = 2

	return cycleFound
}
