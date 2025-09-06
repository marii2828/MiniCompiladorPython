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

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, fmt.Errorf("Stack underflow")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, fmt.Errorf("empty Stack")
	}
	return s.data[len(s.data)-1], nil
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

func (s *Stack[any]) PrintStack() {
	fmt.Println("Stack contents:")
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Printf("  [%d]: %v (type: %T)\n", i, s.data[i], s.data[i])
	}
	fmt.Println("End of Stack")
}
