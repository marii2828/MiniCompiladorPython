package main

import (
	"bufio"
	"fmt"
	"minicomp/filelecture"
	"minicomp/internal"
	"minicomp/logistics"
	"os"
	"strconv"
	"strings"
)

//ANSI escape codes for colors
const (
    Reset  = "\033[0m"
    Blue   = "\033[34m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
)

// Displays a formatted menu and gets the user's choice
func ShowMenu(r *bufio.Reader) int {
    //Clear console 
    logistics.ClearConsole()

    fmt.Println(Blue + "=======================================" + Reset)
    fmt.Println(Blue + "         GOLAND COMPILER MENU          " + Reset)
    fmt.Println(Blue + "=======================================" + Reset)
    fmt.Println(Green + "1. Compile a specific file" + Reset)
    fmt.Println(Green + "2. Compile all files in a directory" + Reset)
    fmt.Println(Green + "3. Exit" + Reset)
    fmt.Println(Yellow + "---------------------------------------" + Reset)
    fmt.Print(Yellow + "Select an option: " + Reset)

    line, _ := r.ReadString('\n')
    n, _ := strconv.Atoi(strings.TrimSpace(line))
    return n
}
	
// showFiles lists the files and prompts for user selection
func showFiles(r *bufio.Reader) {
    path := "./tests"
    
    files := filelecture.ParseFiles(path)

    if len(files) == 0 {
        fmt.Println("No files found in the directory:", path)
        return
    }

    fmt.Println("Files available:")
    for i, f := range files {
        fmt.Printf(Blue + "%d. %s\n" + Reset, i+1, f.Name())
    }

    fmt.Print("\nEnter the number of the file you want to compile: ")
    input, _ := r.ReadString('\n')
    input = strings.TrimSpace(input)
    num, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid number.")
        return
    }
    
    if num < 1 || num > len(files) {
        fmt.Println("Invalid file number. Please try again.")
        return
    }
    
    f := files[num-1]
    fmt.Println("You selected the file:", f.Name())
    
    filelecture.ReadFile(path + "/" + f.Name())
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        option := ShowMenu(reader)
        switch option {
        case 1:
            logistics.ClearConsole()
            fmt.Println("You selected to compile a specific file.")
            showFiles(reader)
            internal.RunVMLoop()  

        case 2:
            logistics.ClearConsole()
            fmt.Println("You selected to compile all files in the directory.")
            path := "./tests"
            files := filelecture.ParseFiles(path)
            if len(files) == 0 {
                fmt.Println("No files found in the directory:", path)
            } else {
                fmt.Println("Executing all files in the directory...")
                for _, f := range files {
                    fmt.Println("\nExecuting file:", f.Name())
                    filelecture.ReadFile(path + "/" + f.Name())
                    internal.RunVMLoop()
                    fmt.Println("\n")
                }
            }

        case 3:
            fmt.Println("Exiting the program. Goodbye!")
            os.Exit(0)
        default:
            fmt.Println("Invalid option. Please try again.")
        }
		
        // Pause before looping back to the menu
        if option >= 1 && option <= 2 {
            fmt.Println("\nPress Enter to continue...")
            reader.ReadString('\n') 
        }
    }
}