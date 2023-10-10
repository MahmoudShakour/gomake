build: clean fmt
    go build -o myapp main.go

clean:
    go clean

run: build
    ./myapp
