package internal

import internal "github.com/MahmoudShakour/gomake.git/internal/parser"



func BuildAdjList(targets []internal.Target) map[int][]int {
	targetNametoTargetId := map[string]int{}
	

	for _, target := range targets {
		targetNametoTargetId[target.Name] = target.Id
	}

	adjList := map[int][]int{}

	for _, target := range targets {
		for _, dependencyName := range target.Dependencies {
			from := target.Id
			to := targetNametoTargetId[dependencyName]

			adjList[from] = append(adjList[from], to)
		}
	}

	return adjList
}


