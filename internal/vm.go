package internal

import (
	"fmt"
	"minicomp/filelecture"
)

func RunVMLoop() {
	list := filelecture.GetInstructions()

	for i, instr := range list {
		v := parseConst(instr.Indexs)
		if v == i {
			switch instr.Instruction {
				//Coloca el valor de la constante en el tope de la pila 
			case "LOAD_CONST":
				val := parseConst(instr.Argument) // convierte "7","4.5","'a'","\"hola\"", "True"/"False"
				fmt.Printf("LOAD_CONST %v (type: %T)\n", val, val)

				//Coloca el valor del contenido de la variable en la pila
			case "LOAD_FAST":
				val := parseConst(instr.Argument)
				fmt.Printf("LOAD_FAST %v (type: %T)\n", val, val)

				//Escribe el contenido del tope de la pila en la variable
			case "STORE_FAST":
				val := parseConst(instr.Argument)
				fmt.Printf("STORE_FAST %v (type: %T)\n", val, val)

				//Carga en el tope de la pila o el valor de la variable o la referencia a la función
			case "LOAD_GLOBAL":
				val := parseConst(instr.Argument)
				fmt.Printf("LOAD_GLOBAL %v (type: %T)\n", val, val)

				//Realiza un salto a la dirección de código de la función
			case "CALL_FUNCTION":
				val := parseConst(instr.Argument)
				fmt.Printf("CALL_FUNCTION %v (type: %T)\n", val, val)

				//Realiza una comparación booleana según el op que reciba
			case "COMPARE_OP":
				val := parseConst(instr.Argument)
				fmt.Printf("COMPARE_OP %v (type: %T)\n", val, val)

				//Realiza una resta de operandos
			case "BINARY_SUBSTRACT":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_SUBSTRACT %v (type: %T)\n", val, val)

				//Realiza una suma de operandos
			case "BINARY_ADD":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_ADD %v (type: %T)\n", val, val)

				//Realiza una multiplicación de operandos
			case "BINARY_MULTIPLY":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_MULTIPLY %v (type: %T)\n", val, val)

				//Realiza una división entera de operandos
			case "BINARY_DIVIDE":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_DIVIDE %v (type: %T)\n", val, val)

				//Realiza un AND lógico
			case "BINARY_AND":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_AND %v (type: %T)\n", val, val)

				//Realiza un OR lógico
			case "BINARY_OR":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_OR %v (type: %T)\n", val, val)

				//Realiza el cálculo del cociente de la división de dos operandos
			case "BINARY_MODULO":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_MODULO %v (type: %T)\n", val, val)

				//Realiza la operación:array[index] = value
			case "STORE_SUBSCR":
				val := parseConst(instr.Argument)
				fmt.Printf("STORE_SUBSCR %v (type: %T)\n", val, val)

				//Carga en el tope de la pila el elemento de un arreglo en el índice indicado
			case "BINARY_SUBSCR":
				val := parseConst(instr.Argument)
				fmt.Printf("BINARY_SUBSCR %v (type: %T)\n", val, val)

				//Salta a la línea de código indicada por “target”
			case "JUMP_ABSOLUTE":
				val := parseConst(instr.Argument)
				fmt.Printf("JUMP_ABSOLUTE %v (type: %T)\n", val, val)

				//Si el tope de la pila es True, salta a “target”
			case "JUMP_IF_TRUE":
				val := parseConst(instr.Argument)
				fmt.Printf("JUMP_IF_TRUE %v (type: %T)\n", val, val)

				//Si el tope de la pila es False, salta a “target”
			case "JUMP_IF_FALSE":
				val := parseConst(instr.Argument)
				fmt.Printf("JUMP_IF_FALSE %v (type: %T)\n", val, val)

				//Construye una lista con “elements” cantidad de elementos
			case "BUILD_LIST":
				val := parseConst(instr.Argument)
				fmt.Printf("BUILD_LIST %v (type: %T)\n", val, val)

				//Termina el programa
			case "END":
				fmt.Println("END")
			default:
				fmt.Printf("Instrucción no reconocida: %s\n", instr.Instruction)
			}
		} else {
			fmt.Printf("WRONG INSTRUCTION INDEX\nThere seems to be a problem with the instruction: %s\n", instr.Instruction)
			// Handle the error (e.g., log it, skip the instruction, etc.) Empty stack, invalid instruction, etc.
		}
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
