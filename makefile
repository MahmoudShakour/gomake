build: clean
    go build -o myapp /home/mahmoudshakour/Workspace/gogogo/HelloWorld/hello.go

clean:
    rm -f myapp

run: build
    ./myapp
