package data

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	integerStack := Stack[int]{
		data: make([]*int, 0),
	}
	stringStack := Stack[string]{
		data: make([]*string, 0),
	}
	integers := make([]int, 0)
	strings := make([]string, 0)
	for i := 0; i < 10; i++ {
		integers = append(integers, i)
		strings = append(strings, fmt.Sprintf("%d", i))
		iInteger := i
		integerStack.Put(&iInteger)
		iString := fmt.Sprintf("%d", i)
		stringStack.Put(&iString)
		if *integerStack.data[i] != integers[i] {
			t.Fail()
		}
		if *stringStack.data[i] != strings[i] {
			t.Fail()
		}
	}
	index := len(integerStack.data) - 1
	for integer, err := integerStack.Pop(); err == nil; integer, err = integerStack.Pop() {
		if *integer != integers[index] {
			t.Fail()
			t.Logf("expected %d but got %d", integers[index], *integer)
		}
		index--
	}
	index = len(stringStack.data) - 1
	for stringValue, err := stringStack.Pop(); err == nil; stringValue, err = stringStack.Pop() {
		if *stringValue != strings[index] {
			t.Fail()
			t.Logf("expected %q but got %q", strings[index], *stringValue)
		}
		index--
	}
}
