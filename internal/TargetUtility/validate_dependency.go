package internal

import (
	"errors"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

var ErrDependencyNotFound = errors.New("given dependency is not found")


func CheckInvalidDependency(targets []internal.Target) error {
	isTarget :=  map[string]int {}
	
	for _,target :=range targets {
		isTarget[target.Name]=1
	}

	for _,target :=range targets {
		for _,dependency :=range target.Dependencies {
			if isTarget[dependency]==0 {
				return ErrDependencyNotFound
			}
		}
	}
	return nil
}