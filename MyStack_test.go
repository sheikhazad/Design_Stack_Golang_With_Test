package mystack

import (
	"fmt"
	"testing"
)

func TestStack(test *testing.T) {

	var ms MyStack
	//Or, ms := make(MyStack, 0)

	//Push int values
	ms.Push(-1)
	ms.Push(2)
	ms.Push(0)
	ms.Push(-4)

	//Stack Size() test
	if ms.Size() != 4 {
		test.Error("Test failed: Expected stack size = 4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected stack size = 4 from [-1, 2, 0, -4]")
	}

	//Stack Top() test
	if ms.Top() != -4 {
		test.Errorf("Test failed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	}

	//Pop() test
	ms.Pop() //[-1, 2, 0]
	if ms.Top() != 0 {
		test.Errorf("Test failed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	} else { //Just for printing
		fmt.Println("Test passed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	}

	//Push string
	fmt.Println("Pushing String:")
	ms.Push("I'm String")
	if ms.Top() != "I'm String" {
		test.Errorf("Test failed: After pushing string, Expected top of stack = 'I'm String' from [-1, 2, 0, I'm String]")
	} else { //Just for printing
		fmt.Println("Test passed: After pushing string, Expected top of stack = 'I'm String' from [-1, 2, 0, I'm String]")
	}

	// Clear() and isEmpty() test
	ms.Clear()
	if ms.isEmpty() != true {
		test.Errorf("Test failed: After Clear(), Stack is expected to be empty.")
	} else { //Just for printing
		fmt.Println("Test passed: After Clear(), Stack is expected to be empty.")
	}
}
