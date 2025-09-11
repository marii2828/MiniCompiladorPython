package internal

import (
	"fmt"
	"minicomp/filelecture"
)

func RunVMLoop() {
	list := filelecture.GetInstructions() // variables globales (p.ej. builtins)

	for ip := 0; ip < len(list); {
		in := list[ip]
		next := ip + 1

		//		filelecture.Instructions.PrintInstructions() // Llama a la función para imprimir las instrucciones
		// Si quieres imprimir las instrucciones, llama a la función así:

		switch in.Instruction {
		case "LOAD_CONST":
			OpLoadConst(&StackVM, in.Argument)
			fmt.Println("Loading constant:", in.Argument)

		case "LOAD_FAST":
			OpLoadFast(&StackVM, LocalVarListVM, in.Argument)
			fmt.Println("Loading local variable:", in.Argument)

		case "STORE_FAST":
			OpStoreFast(&StackVM, LocalVarListVM, in.Argument)
			fmt.Println("Storing in local variable:", in.Argument)

		case "LOAD_GLOBAL":
			// PrintVars(GlobalVarListVM)
			OpLoadGlobal(&StackVM, GlobalVarListVM, in.Argument)
			fmt.Println("Loading global variable:", in.Argument)

		case "CALL_FUNCTION":
			// fmt.Println("Calling function:", in.Argument)
			// PrintVars(GlobalVarListVM)
			// PrintVars(LocalVarListVM)
			// StackVM.PrintStack()
			OpCallFunction(&StackVM, GlobalVarListVM, in.Argument)
			fmt.Println("Calling function:", in.Argument)

		case "COMPARE_OP":
			OpCompare(&StackVM, in.Argument)
			fmt.Println("Comparing with operator:", in.Argument)

		case "BINARY_ADD", "BINARY_SUBSTRACT", "BINARY_MULTIPLY",
			"BINARY_DIVIDE", "BINARY_MODULO":
			OpBinary(&StackVM, in.Instruction)
			fmt.Println("Binary operation:", in.Instruction)

		case "BINARY_AND", "BINARY_OR":
			OpLogical(&StackVM, in.Instruction)
			fmt.Println("Binary logical operation:", in.Instruction)

		case "STORE_SUBSCR":
			OpStoreSubscr(&StackVM)
			fmt.Println("Storing in subscription")

		case "BINARY_SUBSCR":
			OpBinarySubscr(&StackVM)
			fmt.Println("Accessing subscription")

		case "JUMP_ABSOLUTE":
			if t, ok := OpJumpAbsolute(in.Argument); ok {
				next = t
			}
			fmt.Println("Jumping to absolute instruction:", in.Argument)

		case "JUMP_IF_TRUE":
			if t, ok := OpJumpIfTrue(&StackVM, in.Argument); ok {
				next = t
			}
			fmt.Println("Jumping if true to instruction:", in.Argument)

		case "JUMP_IF_FALSE":
			if t, ok := OpJumpIfFalse(&StackVM, in.Argument); ok {
				next = t
			}
			fmt.Println("Jumping if false to instruction:", in.Argument)

		case "BUILD_LIST":
			OpBuildList(&StackVM, in.Argument)
			fmt.Println("Building list with", in.Argument, "elements")

		case "END":
			StackVM.PrintStack()
			PrintVars(LocalVarListVM)

			return

		default:
			panic("unknown opcode: " + in.Instruction)
		}

		ip = next
	}
}

// parseConst converts a string constant to its Go value (int, float64, bool, or string).
func parseConst(param string) interface{} {
	// Try to parse as int
	var i int
	_, err := fmt.Sscanf(param, "%d", &i)
	if err == nil {
		return i
	}
	// Try to parse as float64
	var f float64
	_, err = fmt.Sscanf(param, "%f", &f)
	if err == nil {
		return f
	}
	// Check for boolean
	if param == "True" {
		return true
	}
	if param == "False" {
		return false
	}
	// Remove quotes for string
	if len(param) > 1 && (param[0] == '"' || param[0] == '\'') && param[len(param)-1] == param[0] {
		return param[1 : len(param)-1]
	}
	return param

	//I think type list exists but i have to ask something first
}
