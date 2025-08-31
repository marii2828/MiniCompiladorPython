package internal

import (
	"fmt"
)

// Pushes a value onto the stack
func (vm *VM) Push(value Value) {
	vm.Stack = append(vm.Stack, value)
}

// Pops a value from the stack
func (vm *VM) Pop() (Value, error) {
	if len(vm.Stack) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	value := vm.Stack[len(vm.Stack)-1]
	vm.Stack = vm.Stack[:len(vm.Stack)-1]
	return value, nil
}

//Compares two values with an operator
func Compare(a, b any, op string) (bool, error) {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int: return CmpInts(x, y, op)
		case float64: return CmpFloats(float64(x), y, op)
		}
	case float64:
		switch y := b.(type) {
		case int: return CmpFloats(x, float64(y), op)
		case float64: return CmpFloats(x, y, op)
		}
	case string:
		if ys, ok := b.(string); ok { return CmpStrings(x, ys, op) }
	}
	return false, typeErr("COMPARE_OP "+op, a, b)
}

func CmpInts(a, b int, op string) (bool, error)   { 
	return CmpFloats(float64(a), float64(b), op) 
}
func CmpStrings(a, b, op string) (bool, error)    { 
	switch op { 
		case "==": 
			return a==b, nil; 
		case "!=": 
			return a!=b, nil; 
		case "<": 
			return a<b, nil; 
		case "<=": 
			return a<=b, nil; 
		case ">": 
			return a>b, nil; 
		case ">=": 
			return a>=b, nil 
		}; 
		return false, fmt.Errorf("op string inválido: %s", op) 
}
func CmpFloats(a, b float64, op string) (bool, error) {
	switch op {
	case "==": 
		return a == b, nil
	case "!=": 
		return a != b, nil
	case "<":  
		return a <  b, nil
	case "<=": 
		return a <= b, nil
	case ">":  
		return a >  b, nil
	case ">=": 
		return a >= b, nil
	}
	return false, fmt.Errorf("op inválido: %s", op)
}

func typeErr(op string, a, b any) error {
	return fmt.Errorf("%s tipos incompatibles: %T y %T", op, a, b)
}