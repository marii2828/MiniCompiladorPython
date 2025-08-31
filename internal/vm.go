package internal

import (
	"fmt"
	"minicomp/filelecture"
	"strings"
)

type Value = any 

type VM struct {
	Stack []Value // Stack to hold values during execution
	Vars map[string]Value // Map to hold variable names and their values
	Ip    int // Instruction pointer
	Code  []filelecture.Instructions // Slice to hold the program instructions
}

// RunVMLoop initializes and runs the virtual machine loop
func RunVMLoop() {
	list := filelecture.GetInstructions()

	// Initialize the VM
	vm := &VM{
		Stack: make([]Value, 0),
		Vars:  make(map[string]Value),
		Ip:    0,
		Code:  list,
	}
	if err := vm.run(); err != nil {
		fmt.Println("Error during execution:", err)
	}
}

// run executes the instructions in the VM's code slice
func (vm *VM) run() error {
	// Main execution loop. Executes instructions while the vm.Ip pointer is within the program limits 
	for vm.Ip >= 0 && vm.Ip < len(vm.Code) {
		instr := vm.Code[vm.Ip]
		// Avance por defecto: +1 (si hay salto, se sobrescribe ip)
		next := vm.Ip + 1
        
		switch instr.Instruction {
				//Coloca el valor de la constante en el tope de la pila 
			case "LOAD_CONST":
				val := parseConst(instr.Argument) // convierte "7","4.5","'a'","\"hola\"", "True"/"False"
				fmt.Printf("LOAD_CONST %v (type: %T)\n", val, val)

				vm.Push(val)
				fmt.Println("Pila actual:", vm.Stack)

				
			case "LOAD_FAST":
				//Coloca el valor del contenido de la variable en la pila
				val := parseConst(instr.Argument)
				fmt.Printf("LOAD_FAST %v (type: %T)\n", val, val)

				//Looks for the value asociated with the variable name in the map
				val, ok := vm.Vars[instr.Argument]
				if !ok {
					return fmt.Errorf("LOAD_FAST variable '%s' no definida (ip=%d)", instr.Argument, vm.Ip)
				}
				vm.Push(val)

				
			case "STORE_FAST":
				//Escribe el contenido del tope de la pila en la variable
				val := parseConst(instr.Argument)
				fmt.Printf("STORE_FAST %v (type: %T)\n", val, val)


				if instr.Argument == "" {
					return fmt.Errorf("variable name is empty")
				}
				val, err := vm.Pop()
				if err != nil {
					return err
				}
				//Asings de popped value to the variable in the Vars map 
				vm.Vars[instr.Argument] = val
				fmt.Println("Variables actuales:", vm.Vars)

				
			case "LOAD_GLOBAL":
				//Carga en el tope de la pila o el valor de la variable o la referencia a la función
				val := parseConst(instr.Argument)
				fmt.Printf("LOAD_GLOBAL %v (type: %T)\n", val, val)

				//Looks for the value asociated with the variable name in the map
				val, ok := vm.Vars[instr.Argument]
				if !ok {
					return fmt.Errorf("LOAD_GLOBAL '%s' no definido (ip=%d)", instr.Argument, vm.Ip)
				}
				vm.Push(val)
				fmt.Println("Pila actual:", vm.Stack)

				
			case "CALL_FUNCTION":
				//Realiza un salto a la dirección de código de la función

				
			case "COMPARE_OP":
				//Realiza una comparación booleana según el op que reciba
				val := parseConst(instr.Argument)
				fmt.Printf("COMPARE_OP %v (type: %T)\n", val, val)

				//The argument is the comparison operator as a string
				op := strings.TrimSpace(instr.Argument) // "<", "<=", "==", "!=", ">", ">="
				b, err := vm.Pop(); if err != nil { 
					return err 
				}
				a, err := vm.Pop(); if err != nil { 
					return err 
				}
				res, err := Compare(a, b, op)
				if err != nil { 
					return err 
				}
				vm.Push(res)

				
			case "BINARY_SUBSTRACT":
				//Realiza una resta de operandos

				
			case "BINARY_ADD":
				//Realiza una suma de operandos

				
			case "BINARY_MULTIPLY":
				//Realiza una multiplicación de operandos

				
			case "BINARY_DIVIDE":
				//Realiza una división entera de operandos

				
			case "BINARY_AND":
				//Realiza un AND lógico
				
			case "BINARY_OR":
				//Realiza un OR lógico

				
			case "BINARY_MODULO":
				//Realiza el cálculo del cociente de la división de dos operandos

				
			case "STORE_SUBSCR":
				//Realiza la operación:array[index] = value

				
			case "BINARY_SUBSCR":
				//Carga en el tope de la pila el elemento de un arreglo en el índice indicado

				
			case "JUMP_ABSOLUTE":
				//Salta a la línea de código indicada por “target”

				
			case "JUMP_IF_TRUE":
				//Si el tope de la pila es True, salta a “target”

				
			case "JUMP_IF_FALSE":
				//Si el tope de la pila es False, salta a “target”

				
			case "BUILD_LIST":
				//Construye una lista con “elements” cantidad de elementos

				//Takes the number of elements to pop from the stack and create the list
				val := parseConst(instr.Argument)
				elements := make([]interface{}, val.(int))
				for i := 0; i < val.(int); i++ {
					elem, err := vm.Pop()
					if err != nil {
						return err
					}
					//Reverse order to maintain the original order in the stack
					elements[val.(int)-1-i] = elem
				}
				vm.Push(elements)
				fmt.Println("Pila actual:", vm.Stack)

				//Termina el programa
			case "END":
				return nil

		default:
			return fmt.Errorf("Instrucción no reconocida: %s (ip=%d)", instr.Instruction, vm.Ip)
		}

		vm.Ip = next
	}
	return nil
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




/*
--------PRIMERA VUELTA DE LAS ISNTRUCCIONES----------------

0 LOAD_CONST 0 → stack=[0]
1 STORE_FAST x → vars[x]=0, stack=[]

2..11 LOAD_CONST 0..9 → tras la 11: stack=[0,1,2,3,4,5,6,7,8,9]
12 BUILD_LIST 10 → toma 10 del tope y deja stack=[[0,1,2,3,4,5,6,7,8,9]]
13 STORE_FAST lista → vars[lista]=[0..9], stack=[]

14 LOAD_FAST x → stack=[0]
15 LOAD_CONST 10 → stack=[0,10]
16 COMPARE_OP < → compara 0 < 10 → stack=[true]
17 JUMP_IF_FALSE 37 → como es true, no salta (descarta el tope); sigue a 18. stack=[]

18 LOAD_FAST x → stack=[0]
19 LOAD_CONST 2 → stack=[0,2]
20 BINARY_MODULO → 0 % 2 = 0 → stack=[0]
21 LOAD_CONST 0 → stack=[0,0]
22 COMPARE_OP == → 0 == 0 → stack=[true]
23 JUMP_IF_FALSE 29 → es true, no salta (descarta el tope); sigue a 24. stack=[]

24 LOAD_GLOBAL print → stack=[ref(print)]
25 LOAD_FAST lista → stack=[ref(print), [0..9]]
26 LOAD_FAST x → stack=[ref(print), [0..9], 0]
27 BINARY_SUBSCR → toma index=0, array=[0..9] → empuja array[index]=0
→ stack=[ref(print), 0]
28 CALL_FUNCTION 1 → saca 1 parámetro (0) + la ref de función y llama → imprime 0
→ stack=[]

29 LOAD_FAST x → stack=[0]
30 LOAD_CONST 1 → stack=[0,1]
31 BINARY_ADD → 0+1=1 → stack=[1]
32 STORE_FAST x → vars[x]=1, stack=[]

33 LOAD_FAST x → stack=[1]
34 LOAD_CONST 10 → stack=[1,10]
35 COMPARE_OP < → 1 < 10 → stack=[true]
36 JUMP_IF_TRUE 18 → como es true, salta a ip=18 (descarta el tope). stack=[]
*/

/*
----------SEGUNDA VUELTA (x=1) – resumen rápido------------

18–23: calcula 1 % 2 == 0 → false → salta a 29 (no imprime).

29–32: x = x + 1 → x=2

33–36: x < 10 → true → salta a 18.

----------------TERCERA VUELTA (x=2) ------------------------

18–23: 2 % 2 == 0 → true → 24–28 imprimen lista[2] = 2.

29–36: x=3, compara <10, vuelve a 18.

--------------- ... CONTINUA --------------------------

Para x par (0,2,4,6,8): se imprime 0, 2, 4, 6, 8.

Para x impar (1,3,5,7,9): no imprime.

Cuando x llega a 10:

14–17: 10 < 10 → false → JUMP_IF_FALSE 37 salta a 37.

37 END → termina el programa.
*/