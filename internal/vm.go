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
			fmt.Println("Cargando constante:", in.Argument)

		case "LOAD_FAST":
			OpLoadFast(&StackVM, LocalVarListVM, in.Argument)
			fmt.Println("Cargando variable local:", in.Argument)

		case "STORE_FAST":
			OpStoreFast(&StackVM, LocalVarListVM, in.Argument)
			fmt.Println("Almacenando en variable local:", in.Argument)

		case "LOAD_GLOBAL":
			OpLoadGlobal(&StackVM, GlobalVarListVM, in.Argument)
			fmt.Println("Cargando variable global:", in.Argument)

		case "CALL_FUNCTION":
			OpCallFunction(&StackVM, GlobalVarListVM, in.Argument)
			fmt.Println("Llamando a función:", in.Argument)

		case "COMPARE_OP":
			OpCompare(&StackVM, in.Argument)
			fmt.Println("Comparando con operador:", in.Argument)

		case "BINARY_ADD", "BINARY_SUBSTRACT", "BINARY_MULTIPLY",
			"BINARY_DIVIDE", "BINARY_MODULO":
			OpBinary(&StackVM, in.Instruction)
			fmt.Println("Operación binaria:", in.Instruction)

		case "BINARY_AND", "BINARY_OR":
			OpLogical(&StackVM, in.Instruction)
			fmt.Println("Operación lógica binaria:", in.Instruction)

		case "STORE_SUBSCR":
			OpStoreSubscr(&StackVM)
			fmt.Println("Almacenando en subíndice")

		case "BINARY_SUBSCR":
			OpBinarySubscr(&StackVM)
			fmt.Println("Accediendo a subíndice")

		case "JUMP_ABSOLUTE":
			if t, ok := OpJumpAbsolute(in.Argument); ok {
				next = t
			}
			fmt.Println("Saltando a instrucción absoluta:", in.Argument)

		case "JUMP_IF_TRUE":
			if t, ok := OpJumpIfTrue(&StackVM, in.Argument); ok {
				next = t
			}
			fmt.Println("Saltando si es verdadero a instrucción:", in.Argument)

		case "JUMP_IF_FALSE":
			if t, ok := OpJumpIfFalse(&StackVM, in.Argument); ok {
				next = t
			}
			fmt.Println("Saltando si es falso a instrucción:", in.Argument)

		case "BUILD_LIST":
			OpBuildList(&StackVM, in.Argument)
			fmt.Println("Construyendo lista con", in.Argument, "elementos")

		case "END":
			fmt.Println("\n\nFin del programa")
			StackVM.PrintStack()
			fmt.Println("\n\nVariables locales:")
			PrintVars(LocalVarListVM)
			fmt.Println("\n\nVariables globales:")
			PrintVars(GlobalVarListVM)

			return

		default:
			panic("opcode desconocido: " + in.Instruction)
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
