package singlylinkedlist_test

import (
	singlylinkedlist "godsa/linked-list/singly-linkedlist"
	"reflect"
	"testing"
)

func TestSinglyLinkedListFunctions(t *testing.T) {
	t.Run("TestNewSinglyLinkedList", func(t *testing.T) {
		data := 42
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		if sll == nil {
			t.Error("NewSinglyLinkedList should return a non-nil pointer")
		}
		if sll.GetData() != data {
			t.Errorf("Expected data %d, but got %d", data, sll.GetData())
		}
		if sll.GetNext() != nil {
			t.Error("Expected next to be nil for a new linked list")
		}
	})

	t.Run("TestGetData", func(t *testing.T) {
		data := 42
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		result := sll.GetData()
		if result != data {
			t.Errorf("Expected GetData() to return %d, but got %d", data, result)
		}
	})

	t.Run("TestSetData", func(t *testing.T) {
		data := 42
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		newData := 24
		sll.SetData(newData)
		if sll.GetData() != newData {
			t.Errorf("Expected SetData() to set data to %d, but it's %d", newData, sll.GetData())
		}
	})

	t.Run("TestGetNext", func(t *testing.T) {
		data := 42
		nextData := 24
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		sll.SetNext(singlylinkedlist.NewSinglyLinkedList(nextData))
		result := sll.GetNext()
		if result == nil {
			t.Error("GetNext() should return a non-nil pointer")
		}
		if result.GetData() != nextData {
			t.Errorf("Expected GetNext() to return data %d, but got %d", nextData, result.GetData())
		}
	})

	t.Run("TestSetNext", func(t *testing.T) {
		data := 42
		nextData := 24
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		nextNode := singlylinkedlist.NewSinglyLinkedList(nextData)
		sll.SetNext(nextNode)
		if sll.GetNext() == nil {
			t.Error("SetNext() should set the next node")
		}
		if sll.GetNext().GetData() != nextData {
			t.Errorf("Expected SetNext() to set the next node's data to %d, but it's %d", nextData, sll.GetNext().GetData())
		}
	})

	t.Run("TestGetLength", func(t *testing.T) {
		data := 42
		nextData := 24
		sll := singlylinkedlist.NewSinglyLinkedList(data)
		sll.SetNext(singlylinkedlist.NewSinglyLinkedList(nextData))
		length := singlylinkedlist.GetLength(sll)
		if length != 2 {
			t.Errorf("Expected GetLength() to return 2, but got %d", length)
		}
	})

	t.Run("TestGetLengthEmptyList", func(t *testing.T) {
		sll := singlylinkedlist.NewSinglyLinkedList(42)
		length := singlylinkedlist.GetLength(sll)
		if length != 1 {
			t.Errorf("Expected GetLength() for an empty list to return 1, but got %d", length)
		}
	})

	t.Run("Test Traversing through singly linked list", func(t *testing.T) {
		head := &singlylinkedlist.SinglyLL{}
		head.SetData(1)
		head.SetNext(&singlylinkedlist.SinglyLL{})
		head.GetNext().SetData(2)
		head.GetNext().SetNext(&singlylinkedlist.SinglyLL{})
		head.GetNext().GetNext().SetData(3)
		// Iterate through the linked list and get the result
		result := singlylinkedlist.TraverseSinglyLinkedList(head)

		// Define the expected result
		expected := []int{1, 2, 3}

		// Check if the result matches the expected value
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}

	})
}

func TestInsertAtBeginning(t *testing.T) {
	t.Run("Insert in empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		// Test inserting a node into an empty list
		data1 := 25
		singlylinkedlist.InsertAtBeginning(&head, data1)

		if head == nil {
			t.Error("Expected head to be non-nil after inserting into an empty list")
		}
		if head.GetData() != data1 {
			t.Errorf("Expected data in the head to be %d, but got %d", data1, head.GetData())
		}

		if head.GetNext() != nil {
			t.Error("Expected next to be nil after inserting into an empty list")
		}
	})

	t.Run("Insert into a non-empty list", func(t *testing.T) {
		data1 := 26
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 25
		singlylinkedlist.InsertAtBeginning(&head, data2)

		if head == nil {
			t.Error("Expected head to be non-nil after inserting into a non-empty list")
		}
		if head.GetData() != data2 {
			t.Errorf("Expected data in the head to be %d, but got %d", data2, head.GetData())
		}
		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after inserting into a non-empty list")
		}
		if head.GetNext().GetData() != data1 {
			t.Errorf("Expected the next node data to be %d, but got %d", data1, head.GetNext().GetData())
		}
	})
}

func TestInsertAtPosition(t *testing.T) {
	t.Run("Insert at position 0 in an empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		data := 42
		position := 0
		err := singlylinkedlist.InsertAtPosition(&head, data, position)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if head == nil {
			t.Error("Expected head to be non-nil after inserting at position 0 in an empty list")
		}
		if head.GetData() != data {
			t.Errorf("Expected data in the head to be %d, but got %d", data, head.GetData())
		}
		if head.GetNext() != nil {
			t.Error("Expected next to benil after inserting at position 0 in an empty list")
		}
	})

	t.Run("Insert at a valid position in a non-empty list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		position := 1
		err := singlylinkedlist.InsertAtPosition(&head, data2, position)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after inserting at a valid position in a non-empty list")
		}
		if head.GetNext().GetData() != data2 {
			t.Errorf("Expected the next node data to be %d, but got %d", data2, head.GetNext().GetData())
		}
	})

	t.Run("Insert at a valid position in a non-empty list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		position := 1
		err := singlylinkedlist.InsertAtPosition(&head, data2, position)
		data3 := 24
		position = 2
		err = singlylinkedlist.InsertAtPosition(&head, data3, position)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after inserting at a valid position in a non-empty list")
		}
		if head.GetNext().GetData() != data2 {
			t.Errorf("Expected the next node data to be %d, but got %d", data2, head.GetNext().GetData())
		}
	})

	t.Run("Insert at an invalid position in a non-empty list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		position := 2
		err := singlylinkedlist.InsertAtPosition(&head, data2, position)

		if err == nil {
			t.Error("Expected an error for inserting at an invalid position in a non-empty list")
		}
		if err.Error() != "position out of range" {
			t.Errorf("Expected error message 'position out of range', but got: %v", err)
		}
	})

	t.Run("Insert at a negative position", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		data := 42
		position := -1
		err := singlylinkedlist.InsertAtPosition(&head, data, position)

		if err == nil {
			t.Error("Expected an error for inserting at a negative position")
		}
		if err.Error() != "invalid position" {
			t.Errorf("Expected error message 'invalid position', but got: %v", err)
		}
	})
}

func TestInsertAtEnd(t *testing.T) {
	t.Run("Insert into an empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		data := 42
		err := singlylinkedlist.InsertAtEnd(&head, data)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head == nil {
			t.Error("Expected head to be non-nil after inserting into an empty list")
		}
		if head.GetData() != data {
			t.Errorf("Expected data in the head to be %d, but got %d", data, head.GetData())
		}
		if head.GetNext() != nil {
			t.Error("Expected next to be nil after inserting into an empty list")
		}
	})

	t.Run("Insert into a non-empty list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		err := singlylinkedlist.InsertAtEnd(&head, data2)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		current := head
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		if current.GetData() != data2 {
			t.Errorf("Expected the last node data to be %d, but got %d", data2, current.GetData())
		}
	})

	t.Run("Insert into a non-empty list of length > 2", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		err := singlylinkedlist.InsertAtEnd(&head, data2)
		data3 := 25
		err = singlylinkedlist.InsertAtEnd(&head, data3)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		current := head
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		if current.GetData() != data3 {
			t.Errorf("Expected the last node data to be %d, but got %d", data2, current.GetData())
		}
	})
}

func TestDeleteFirstNode(t *testing.T) {
	t.Run("Delete from an empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		err := singlylinkedlist.DeleteFirstNode(&head)

		if err == nil {
			t.Error("Expected an error for deleting from an empty list")
		}
		if err.Error() != "nothing to delete" {
			t.Errorf("Expected error message 'nothing to delete', but got: %v", err)
		}
	})

	t.Run("Delete from a single-node list", func(t *testing.T) {
		data := 42
		head := singlylinkedlist.NewSinglyLinkedList(data)
		err := singlylinkedlist.DeleteFirstNode(&head)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head != nil {
			t.Error("Expected head to be nil after deleting from a single-node list")
		}
	})

	t.Run("Delete from a multi-node list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtBeginning(&head, data2)
		err := singlylinkedlist.DeleteFirstNode(&head)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head.GetData() != data1 {
			t.Errorf("Expected the new head data to be %d, but got %d", data1, head.GetData())
		}
	})
}

func TestDeleteLastNode(t *testing.T) {
	t.Run("Delete from an empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		err := singlylinkedlist.DeleteLastNode(&head)

		if err == nil {
			t.Error("Expected an error for deleting from an empty list")
		}
		if err.Error() != "nothing to delete" {
			t.Errorf("Expected error message 'nothing to delete', but got: %v", err)
		}
	})

	t.Run("Delete from a single-node list", func(t *testing.T) {
		data := 42
		head := singlylinkedlist.NewSinglyLinkedList(data)
		err := singlylinkedlist.DeleteLastNode(&head)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head != nil {
			t.Error("Expected head to be nil after deleting from a single-node list")
		}
	})

	t.Run("Delete from a 2 node list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtEnd(&head, data2)
		err := singlylinkedlist.DeleteLastNode(&head)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if head.GetNext() != nil {
			t.Error("Expected next to be non-nil after deleting from a 2 list")
		}
		if head.GetData() != data1 {
			t.Errorf("Expected the new last node data to be %d, but got %d", data1, head.GetData())
		}
	})

	t.Run("Delete from a multi node list", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtEnd(&head, data2)
		data3 := 25
		singlylinkedlist.InsertAtEnd(&head, data3)
		err := singlylinkedlist.DeleteLastNode(&head)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after deleting from a 2 list")
		}
		if head.GetNext().GetData() != data2 {
			t.Errorf("Expected the new last node data to be %d, but got %d", data2, head.GetNext().GetData())
		}
	})
}

func TestDeleteIntermediateNode(t *testing.T) {
	t.Run("Delete from an empty list", func(t *testing.T) {
		var head *singlylinkedlist.SinglyLL
		value := 42
		err := singlylinkedlist.DeleteIntermediateNode(&head, value)

		if err == nil {
			t.Error("Expected an error for deleting from an empty list")
		}
		if err.Error() != "nothing to delete" {
			t.Errorf("Expected error message 'nothing to delete', but got: %v", err)
		}
	})

	t.Run("Delete the first node", func(t *testing.T) {
		data := 42
		head := singlylinkedlist.NewSinglyLinkedList(data)
		value := 42
		err := singlylinkedlist.DeleteIntermediateNode(&head, value)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head != nil {
			t.Error("Expected head to be nil after deleting the first node")
		}
	})

	t.Run("Delete a middle node", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtEnd(&head, data2)
		data3 := 25
		singlylinkedlist.InsertAtEnd(&head, data3)
		value := 42
		err := singlylinkedlist.DeleteIntermediateNode(&head, value)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after deleting a middle node")
		}
		if head.GetNext().GetData() != data3 {
			t.Errorf("Expected the next node data to be %d, but got %d", data3, head.GetNext().GetData())
		}
	})

	t.Run("Delete a middle node", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtEnd(&head, data2)
		data3 := 25
		singlylinkedlist.InsertAtEnd(&head, data3)
		value := 24
		err := singlylinkedlist.DeleteIntermediateNode(&head, value)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if head.GetNext() == nil {
			t.Error("Expected next to be non-nil after deleting a middle node")
		}
		if head.GetNext().GetData() != data3 {
			t.Errorf("Expected the next node data to be %d, but got %d", data3, head.GetNext().GetData())
		}
	})

	t.Run("Delete a non-existing value", func(t *testing.T) {
		data1 := 42
		head := singlylinkedlist.NewSinglyLinkedList(data1)
		data2 := 24
		singlylinkedlist.InsertAtEnd(&head, data2)
		value := 12
		err := singlylinkedlist.DeleteIntermediateNode(&head, value)

		if err == nil {
			t.Error("Expected an error for deleting a non-existing value")
		}
		if err.Error() != "not found" {
			t.Errorf("Expected error message 'not found', but got: %v", err)
		}
	})
}
