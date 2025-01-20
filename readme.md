# Creating a DLL in GO

This repository contains a simple example of creating a DLL library in Go for further integration into other programs.

# To create a DLL, follow these steps:

## 1) Create a Go Project
    go mod init Go_DLL_Create_Example

## 2) Insert the following code:
    package main

    import "C"
    import "fmt"

    //export sayHello
    func sayHello() {
        fmt.Println("Hello world!")
    }

    func main() {}

### Important
func main - must be empty
Above the methods that should be accessible for calling, there must be a comment 
    
    //export {method name}

## 3) Prepare the Environment
Check Go environment variables
    
    go env

By default, Go has the setting CGO_ENABLED=0, it needs to be set to 1 (can be checked with the command)
    
    go env -w CGO_ENABLED=1

Install gcc
    
    choco install mingw -y

* How to install choco, I think you can find out yourself :)

## 4) Run the DLL creation command
    go build -buildmode=c-shared -o example_from_go.dll main.go

## 5) Check the DLL functionality
### Check 1:
Execute the following commands in the terminal:
        
        Rundll32.exe .\example_from_go.dll,sayHello
If no error occurs, then everything is okay.
fmt.Println("Hello world!") - will not be displayed on the screen!!!

### Check 2:
Create a file checkDll.go and insert the following code
        
    package main

    import "C"
    import (
        "fmt"
        "syscall"
    )

    func main() {
        dll := syscall.MustLoadDLL("example_from_go.dll")
        proc := dll.MustFindProc("sayHello")
        proc.Call()
        fmt.Println("Function executed successfully.")
    }

Here, it should output
    
    Hello world!
    Function executed successfully.
