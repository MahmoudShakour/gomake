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
		want := target{name: "target1", dependencies: []string{"dep1", "dep2", "dep3"}, command: "this is a command"}

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
		want := []target{
			{name: "build", dependencies: []string{}, command: "go build -o bin/main main.go"},
			{name: "run", dependencies: []string{}, command: "go run main.go"},
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
		want := []target{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
