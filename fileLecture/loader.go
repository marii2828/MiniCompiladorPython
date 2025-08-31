package filelecture

import (
    "fmt"
    "os"
    "strings"
	"log"
)

//struct containing the main parts of the instructions on the txt file
type instructions struct {
    indexs  string
    instruction  string
    argument     string
}

//type alias fo a slice of the struct "instructions" 
type programInstructions []instructions

//instance of the slice
var instructionsList programInstructions

//method to process the content of the file and add instructions to the slice
func (p *programInstructions) addInstruction(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		parts := strings.Fields(trimmedLine)
		if len(parts) >= 3 {
			newInstruction := instructions{
				indexs:     parts[0],
				instruction: parts[1],
				argument:    strings.Join(parts[2:], " "),
			}
			*p = append(*p, newInstruction)
		} else {
			fmt.Printf("Advertencia: La línea '%s' no tiene el formato esperado y será omitida.\n", trimmedLine)
		}
	}
	fmt.Println("Procesamiento terminado.")
}

func ReadFile(filePath string) {
	//Check if a file path is provided as a command-line argument
	if filePath == "" {
		fmt.Println("Uso: go run tu_programa.go <nombre_del_archivo>")
		return
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Clear previous instructions and prints them 
	instructionsList = nil
	instructionsList.addInstruction(string(fileContent))
	for _, inst := range instructionsList {
		fmt.Printf("Índice: %s, Instrucción: %s, Argumento: %s\n", inst.indexs, inst.instruction, inst.argument)
	}
}

func ParseFiles(filePath string) []os.DirEntry{

	files, err := os.ReadDir(filePath)
			if err != nil {
				log.Fatal("Error finding tests:", err)
			}

			fmt.Println("Files founded in tests:")
			for i, f := range files {
				if f.IsDir() {
					fmt.Println("first dir", f.Name())
				} else {
					fmt.Println(i+1, " File:", f.Name())
				}
			}
	return files
}