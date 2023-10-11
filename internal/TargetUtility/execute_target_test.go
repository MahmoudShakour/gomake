package internal

import (
	"reflect"
	"testing"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)


func TestBuildTopoSort(t *testing.T) {
	
	t.Run("target with no dependencies",func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"},
			{Id: 1, Name: "clean", Dependencies: []string{}, Command: "go clean"},
			{Id: 2, Name: "run", Dependencies: []string{"build"}, Command: "./myapp"},
		}
		got:=BuildTopoSort(1,make([]int, 3),BuildAdjList(targets))
		want:=[]int {1}

		if  !reflect.DeepEqual(got,want) {
			t.Errorf("got %v want %v",got,want)
		}
	})

	t.Run("target with linear dependencies",func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"},
			{Id: 1, Name: "clean", Dependencies: []string{}, Command: "go clean"},
			{Id: 2, Name: "run", Dependencies: []string{"build"}, Command: "./myapp"},
		}
		got:=BuildTopoSort(2,make([]int, 3),BuildAdjList(targets))
		want:=[]int {1,0,2}

		if  !reflect.DeepEqual(got,want) {
			t.Errorf("got %v want %v",got,want)
		}
	})	
}