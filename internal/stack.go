package internal

import "fmt"

type Stack[T any] struct {
	data []T
}

var StackVM Stack[any] = *NewStack[any]()

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Sentinel errors para detectar fallos sin depender de strings
var ErrStackUnderflow = fmt.Errorf("Stack underflow")
var ErrEmptyStack = fmt.Errorf("empty Stack")

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, ErrStackUnderflow
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

// Ensure valida que haya al menos n items en la pila
func (s *Stack[T]) Ensure(n int) error {
	if len(s.data) < n {
		return ErrStackUnderflow
	}
	return nil
}

// PopN extrae n elementos de golpe y los devuelve en orden lógico
func (s *Stack[T]) PopN(n int) ([]T, error) {
	if n < 0 || len(s.data) < n {
		var zero []T
		return zero, ErrStackUnderflow
	}
	start := len(s.data) - n
	out := make([]T, n)
	copy(out, s.data[start:])
	s.data = s.data[:start]
	return out, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Reset() {
	s.data = s.data[:0]
}

// Clear libera toda la memoria de la pila
func (s *Stack[T]) Clear() {
	s.data = nil
}

func (s *Stack[any]) PrintStack() {
	fmt.Println("\n\n----------------STACK CONTENT (RESULT)----------------")
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Printf("  [%d]: %v (type: %T)\n", i, s.data[i], s.data[i])
	}
	fmt.Println("\n----------------END OF STACK----------------")
}
