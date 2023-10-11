build: clean 
    go build -o myapp main.go

clean:
    go clean

run: build
    ./myapp
