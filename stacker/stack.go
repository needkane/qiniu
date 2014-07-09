package stacker

import (
	"errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

//return capacity(> len)
func (stack Stack) Cap() int {
	return cap(stack)
}
func IsEmpty(stack Stack) bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

//stack是引用的地址，这里*stack是地址指向的值
func (stack *Stack) Push(x interface{}) {

	*stack = append(*stack, x)
}
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't top an empty stack")
	}
	return stack[len(stack)-1], nil
}
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't pop an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}
