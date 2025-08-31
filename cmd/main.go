package main

import (
	
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"../logistics"
	"time"
	"../filelecture"
)

func ShowMenu(r *bufio.Reader) int {
    fmt.Println("Welcome to the Goland compiler")
    fmt.Println("1. Compile a specific file")
    fmt.Println("2. Compile all files in a directory")
    fmt.Println("3. Exit")
    fmt.Print("Select an option: ")
    line, _ := r.ReadString('\n')
    n, _ := strconv.Atoi(strings.TrimSpace(line))
    return n
}
	

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//runtime.ClearConsole()
		option := ShowMenu(reader)
		switch option {
		case 1:
			// Clear the console and show the files in the txt-files directory
			time.Sleep(1 * time.Second)
			// Clear the console
			logistics.ClearConsole()
			fmt.Println("You selected to compile a specific file.")

			// Show the files in the txt-files directory
			path := "./tests"

			// Read the directory
			files := filelecture.ParseFiles(path)

			// Ask the user to select a file
			fmt.Print("Select the file you want to compile: ")
			fmt.Print("Enter the number of the file you want to compile: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(num, " is not a valid number.")
				return
			}
			if num < 0 || num > len(files) {
				fmt.Println("Invalid number. Please try again.")
			} else {
				f := files[num-1]
				fmt.Println("You selected the file:", f.Name())

				// Reads the file in runtime
				filelecture.ReadFile(path + "/" + f.Name())
			}
			
			time.Sleep(3 * time.Second)
			logistics.ClearConsole()

		case 2:
			fmt.Println("You selected to compile all files in the directory.")
			// Show the files in the txt-files directory
			path := "./txt-files"
			files := filelecture.ParseFiles(path)

			fmt.Println("Executing all files in the directory...")
			for _, f := range files {
				fmt.Println("\nExecuting file:", f.Name())
				filelecture.ReadFile(path + "/" + f.Name())
				fmt.Println("\n")

			}

			time.Sleep(3 * time.Second)

		case 3:
			fmt.Println("Exiting the program. Goodbye!")
			os.Exit(0)

		case 4:
			fmt.Println("You selected to test the stack.")
			//types.TesterStackNode()

		case 5:
			fmt.Println("You selected to test the empty interface.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
