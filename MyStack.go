/************************************************************************
* This is a non-thread safe but exception safe Stack.
* Thread safe Stack is implemented in other repository :
* https://github.com/sheikhazad/Minimum-Stack-Design_Thread-Safe_With-Test_Golang
*************************************************************************/
package mystack

import (
	"errors"
	"fmt"
)

type MyStack []interface{}

//Push element into the stack
func (this *MyStack) Push(element interface{}) {
	*this = append(*this, element)
}

//Pop last/top element from the stack
func (this *MyStack) Pop() {
	size := len(*this)
	if size == 0 {
		fmt.Println("Stack is empty.")
		return
	}

	*this = (*this)[:size-1]
}

//get last/top element from stack
func (this MyStack) Top() interface{} {
	size := len(this)
	if size == 0 {
		return errors.New("Stack is empty.")
	}

	return this[size-1]
}

func (this MyStack) isEmpty() bool {
	return len(this) == 0
}

func (this *MyStack) Clear() {
	*this = nil
}

func (this MyStack) Size() int {
	return len(this)
}
