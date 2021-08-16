package stack

import (
	"container/list"
	"reflect"
	"testing"
)

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T)  {
	var newStack Stack

	newStack.push(1)
	newStack.push(2)

	t.Run("Stack With Linked List", func(t *testing.T) {
		t.Run("Stack Push", func(t *testing.T) {
			result := newStack.show()
			expected := []int{2,1}
			for x := range result{
				if result[x] != expected[x]{
					t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
				}
			}
		})


		t.Run("Stack isEmpty", func(t *testing.T) {
			if newStack.isEmpty(){
				t.Error("Stack isEmpty is returned true but expected false", newStack.isEmpty())
			}

		})


		t.Run("Stack Length", func(t *testing.T) {
			if newStack.len() != 2{
				t.Error("Stack Length should be 2 but got", newStack.len())
			}
		})

		newStack.pop()
		pop := newStack.pop()

		t.Run("Stack Pop", func(t *testing.T) {
			if pop != 1{
				t.Error("Stack Pop should return 1 but is returned", pop)
			}

		})

		newStack.push(52)
		newStack.push(23)
		newStack.push(99)
		t.Run("Stack Peak", func(t *testing.T) {
			if   newStack.peak() != 99{
				t.Error("Stack Peak should return 99 but got ", newStack.peak())
			}
		})

	})
}


// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T)  {
	t.Run("Stack With Array", func(t *testing.T) {

		stackPush(2)
		stackPush(3)

		t.Run("Stack Push", func(t *testing.T) {
			if !reflect.DeepEqual([]int{3,2}, stackArray){
				t.Errorf("Stack Push is not work we expected %v but got %v", []int{3,2}, stackArray)
			}
		})

		pop := stackPop()

		t.Run("Stack Pop", func(t *testing.T) {

			if stackLength() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})

		stackPush(2)
		stackPush(83)

		t.Run("Stack Peak", func(t *testing.T) {

			if stackPeak() != 83 {
				t.Errorf("Stack Peak is not work we expected %v but got %v", 83, stackPeak())
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if stackLength() != 3{
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackLength())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if stackEmpty() == true{
				t.Errorf("Stack Emtpy is not work we expected %v but got %v", false, stackEmpty())
			}

			stackPop()
			stackPop()
			stackPop()

			if stackEmpty() == false{
				t.Errorf("Stack Emtpy is not work we expected %v but got %v", true, stackEmpty())
			}
		})
	})
}






// TestStackLinkedListWithList for testing Stack with Container/List Library
func TestStackLinkedListWithList(t *testing.T)  {
	stackList := &SList{
		stack: list.New(),
	}


	t.Run("Stack Push", func(t *testing.T) {

		stackList.Push(2)
		stackList.Push(3)

		if stackList.Length() != 2{
			t.Errorf("Stack Push is not work we expected %v but got %v", 2, stackList.Length())
		}
	})


	t.Run("Stack Pop", func(t *testing.T) {
		pop, _ := stackList.Pop()

		if stackList.Length() == 1 && pop != 3 {
			t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
		}
	})



	t.Run("Stack Peak", func(t *testing.T) {

		stackList.Push(2)
		stackList.Push(83)
		peak, _ := stackList.Peak()
		if peak != 83 {
			t.Errorf("Stack Peak is not work we expected %v but got %v", 83, peak)
		}
	})

	t.Run("Stack Length", func(t *testing.T) {
		if stackList.Length() != 3{
			t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackList.Length())
		}
	})

	t.Run("Stack Empty", func(t *testing.T) {
		if stackList.Empty() == true{
			t.Errorf("Stack Emtpy is not work we expected %v but got %v", false, stackList.Empty())
		}

		stackList.Pop()
		stackList.Pop()
		stackList.Pop()

		if stackList.Empty() == false{
			t.Errorf("Stack Emtpy is not work we expected %v but got %v", true, stackList.Empty())
		}
	})
}