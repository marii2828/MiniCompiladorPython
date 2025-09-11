package filelecture

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// struct containing the main parts of the instructions on the txt file
type Instructions struct {
	Indexs      string
	Instruction string
	Argument    string
}

// type alias fo a slice of the struct "instructions"
type ProgramInstructions []Instructions

func PrintInstructions(InstructionsList ProgramInstructions) {
	fmt.Println("\n\n---------------- INSTRUCTIONS LOADED ----------------")
	for _, instr := range InstructionsList {
		fmt.Printf("Index: %s, Instruction: %s, Argument: %s\n", instr.Indexs, instr.Instruction, instr.Argument)
	}
	fmt.Println("\n\n------------------EXECUTION------------------------\n")
}

// instance of the slice
var InstructionsList ProgramInstructions

// method to process the content of the file and add instructions to the slice
func (p *ProgramInstructions) addInstruction(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		parts := strings.Fields(trimmedLine)

		if len(parts) >= 2 { // <— antes pedías >= 3
			arg := ""
			if len(parts) > 2 {
				arg = strings.Join(parts[2:], " ")
			}
			newInstruction := Instructions{
				Indexs:      parts[0],
				Instruction: parts[1],
				Argument:    arg,
			}
			*p = append(*p, newInstruction)
		} else {
			fmt.Printf("Warning: Line '%s' does not have the expected format and will be omitted.\n", trimmedLine)
		}
	}
	PrintInstructions(*p)

}

func ReadFile(filePath string) {
	//Check if a file path is provided as a command-line argument
	if filePath == "" {
		fmt.Println("Uso: go run tu_programa.go <nombre_del_archivo>")
		return
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Clear previous instructions and prints them
	InstructionsList = nil
	InstructionsList.addInstruction(string(fileContent))
}

func ParseFiles(filePath string) []os.DirEntry {

	files, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatal("Error finding tests:", err)
	}

	return files
}

// Returns the instructions list
func GetInstructions() ProgramInstructions {
	return InstructionsList
}
