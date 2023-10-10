package internal

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestParseMakeFile(t *testing.T) {

	t.Run("one target", func(t *testing.T) {

		const body = `target1: dep1 dep2 dep3
    this is a command`
		const filename = "makefile"

		fs := fstest.MapFS{
			filename: {Data: []byte(body)},
		}

		got := ParseMakeFile(fs, filename)
		want := Target{Name: "target1", Dependencies: []string{"dep1", "dep2", "dep3"}, Command: "this is a command"}

		if !reflect.DeepEqual(got[0], want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("multiple targets", func(t *testing.T) {

		const body = `build:
    go build -o bin/main main.go


run:
    go run main.go`
		const filename = "makefile"

		fs := fstest.MapFS{
			filename: {Data: []byte(body)},
		}

		got := ParseMakeFile(fs, filename)
		want := []Target{
			{Name: "build", Dependencies: []string{}, Command: "go build -o bin/main main.go"},
			{Name: "run", Dependencies: []string{}, Command: "go run main.go"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("no targets", func(t *testing.T) {

		const body = ``
		const filename = "makefile"

		fs := fstest.MapFS{
			filename: {Data: []byte(body)},
		}

		got := ParseMakeFile(fs, filename)
		want := []Target{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
