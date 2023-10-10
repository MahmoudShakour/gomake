package internal

import (
	"testing"

	internal "github.com/MahmoudShakour/gomake.git/internal/parser"
)

func TestCheckInvalidDependency(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		var targets = []internal.Target{{Name: "build", Dependencies: []string{"clean"}, Command: "go build -o myapp main.go"}, {Name: "clean", Dependencies: []string{}, Command: "go clean"}}
		got := CheckInvalidDependency(targets)

		if got != nil {
			t.Errorf("got an erorr but the case is valid")
		}
	})

	t.Run("invalid case", func(t *testing.T) {
		var targets = []internal.Target{{Name: "build", Dependencies: []string{"clean","invalid dependency"}, Command: "go build -o myapp main.go"}, {Name: "clean", Dependencies: []string{}, Command: "go clean"}}
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
