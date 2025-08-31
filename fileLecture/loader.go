package filelecture

import (
    "fmt"
    "os"
    "strings"
	"log"
)

//struct containing the main parts of the instructions on the txt file
type Instructions struct {
    Indexs     string
    Instruction string
    Argument    string
}

//type alias fo a slice of the struct "instructions" 
type ProgramInstructions []Instructions

//instance of the slice
var InstructionsList ProgramInstructions

//method to process the content of the file and add instructions to the slice
func (p *ProgramInstructions) addInstruction(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		parts := strings.Fields(trimmedLine)
		if len(parts) >= 3 {
			newInstruction := Instructions{
				Indexs:     parts[0],
				Instruction: parts[1],
				Argument:    strings.Join(parts[2:], " "),
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
	InstructionsList = nil
	InstructionsList.addInstruction(string(fileContent))
}

func ParseFiles(filePath string) []os.DirEntry{

	files, err := os.ReadDir(filePath)
			if err != nil {
				log.Fatal("Error finding tests:", err)
			}

	return files
}

//Returns the instructions list
func GetInstructions() ProgramInstructions {
	return InstructionsList
}