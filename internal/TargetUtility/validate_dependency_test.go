package internal

import (
	"testing"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

func TestCheckInvalidDependency(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		var targets = []internal.Target{
			{Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"}, 
			{Name: "clean", Dependencies: []string{}, Command: "go clean"},
		}
		got := CheckInvalidDependency(targets)

		if got != nil {
			t.Errorf("got an erorr but the case is valid")
		}
	})

	t.Run("invalid case", func(t *testing.T) {
		var targets = []internal.Target{
			{Name: "build", Dependencies: []string{"clean", "invalid dependency"},Command: "go build -o myapp main.go"}, 
			{Name: "clean", Dependencies: []string{}, Command: "go clean"},
		}
		got := CheckInvalidDependency(targets)

		if got != ErrDependencyNotFound {
			t.Errorf("expected error but didn't get one")
		}
	})

	t.Run("empty input target", func(t *testing.T) {
		var targets = []internal.Target{}
		got := CheckInvalidDependency(targets)

		if got != nil {
			t.Errorf("got an erorr but the case is valid")
		}
	})
}

func TestChechCyclicDependency(t *testing.T) {
	t.Run("simple non-cyclic dependency", func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"},
			{Id: 1, Name: "clean", Dependencies: []string{}, Command: "go clean"},
			{Id: 2, Name: "run", Dependencies: []string{"build"}, Command: "./myapp"},
		}

		got := CheckCyclicDependency(targets)

		if got != nil {
			t.Errorf("expected no error but got: %q",got)
		}
	})

	t.Run("simple cyclic dependency", func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"},
			{Id: 1, Name: "clean", Dependencies: []string{"run"}, Command: "go clean"},
			{Id: 2, Name: "run", Dependencies: []string{"build"}, Command: "./myapp"},
		}
		got := CheckCyclicDependency(targets)

		if got != ErrCyclicDependencyFound {
			t.Errorf("expected error but didn't get one")
		}
	})

	t.Run("empty target input", func(t *testing.T) {
		var targets = []internal.Target{}

		got := CheckCyclicDependency(targets)

		if got != nil {
			t.Errorf("expected no error but got one")
		}
	})

	t.Run("diamond shape non-cyclic dependency", func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "a", Dependencies: []string{"b","d"}, Command: ""},
			{Id: 1, Name: "b", Dependencies: []string{"c"}, Command: ""},
			{Id: 3, Name: "d", Dependencies: []string{"c"}, Command: ""},
			{Id: 2, Name: "c", Dependencies: []string{}, Command: ""},
		}

		got := CheckCyclicDependency(targets)

		if got != nil {
			t.Errorf("expected no error but got: %q",got)
		}
	})

	t.Run("complex cyclic dependency", func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "a", Dependencies: []string{"b","c"}, Command: ""},
			{Id: 1, Name: "b", Dependencies: []string{"c"}, Command: ""},
			{Id: 2, Name: "c", Dependencies: []string{"d","e"}, Command: ""},
			{Id: 3, Name: "d", Dependencies: []string{}, Command: ""},
			{Id: 4, Name: "e", Dependencies: []string{"a"}, Command: ""},
		}
		got := CheckCyclicDependency(targets)

		if got != ErrCyclicDependencyFound {
			t.Errorf("expected error but didn't get one")
		}
	})

	t.Run("forest cyclic dependency", func(t *testing.T) {
		var targets = []internal.Target{
			{Id: 0, Name: "a", Dependencies: []string{"b","c"}, Command: ""},
			{Id: 1, Name: "b", Dependencies: []string{"c"}, Command: ""},
			{Id: 2, Name: "d", Dependencies: []string{"e"}, Command: ""},
			{Id: 3, Name: "e", Dependencies: []string{"f"}, Command: ""},
			{Id: 4, Name: "f", Dependencies: []string{"d"}, Command: ""},
		}
		got := CheckCyclicDependency(targets)

		if got != ErrCyclicDependencyFound {
			t.Errorf("expected error but didn't get one")
		}
	})

}
