package internal

import (
	"fmt"
	"os/exec"
	"strings"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)


func buildTopoSort(node int,visited []int,adjList map[int][]int) []int{
	visited[node]=1
	topoSort:=make([]int,0)
	for _,to :=range adjList[node] {
		if visited[to]==0{
			topoSort=append(topoSort,buildTopoSort(to,visited,adjList)...)
		}
	}
	topoSort=append(topoSort, node)
	return topoSort
}

func ExecuteTarget(targets []internal.Target, targetId int,numberOfTargets int) error {
	err := CheckInvalidDependency(targets)
	if err != nil {
		return err
	}
	err = CheckCyclicDependency(targets)
	if err != nil {
		return err
	}
	
	targetIdToTargetCommand:=map[int]string {}
	for _,target:=range targets {
		targetIdToTargetCommand[target.Id]=target.Command
	}

	adjList:=BuildAdjList(targets)
	visited:=make([]int,numberOfTargets)
	topoSort:=buildTopoSort(targetId,visited,adjList)

	
	for _,targetId:=range topoSort {
		targetCommand:=targetIdToTargetCommand[targetId]
		output,err:=executeCommand(targetCommand)
		if err!=nil {
			return err	
		}
		fmt.Println(output)
	}
	return nil
}


func executeCommand(command string) (string,error){
	splittedCommand:=strings.Split(command, " ")

	cmd:=exec.Command(splittedCommand[0],splittedCommand[1:]...)
	output,err:=cmd.CombinedOutput()

	return string(output),err
}