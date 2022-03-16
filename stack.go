package data

import "fmt"

type Stack[datatype any] struct {
	data []*datatype
}

func NewStack[datatype any]() Stack[datatype] {
	return Stack[datatype]{
		data: make([]*datatype, 0),
	}
}

// Put adds newData to the top of the stack.
func (stack *Stack[datatype]) Put(newData *datatype) {
	stack.data = append(stack.data, newData)
}

// Pop will remove and return the data that is on top of the stack.
// It will return non-nil error if the stack was empty.
func (stack *Stack[datatype]) Pop() (*datatype, error) {
	if len(stack.data) == 0 {
		return nil, fmt.Errorf("empty")
	}
	data := stack.data[len(stack.data)-1]
	stack.data[len(stack.data)-1] = nil
	stack.data = stack.data[:len(stack.data)-1]
	return data, nil
}
