package singlylinkedlist

import (
	"errors"
)

type SinglyLL struct {
	data int
	next *SinglyLL
}

func NewSinglyLinkedList(data int) *SinglyLL {
	return &SinglyLL{data: data}
}

func (sll *SinglyLL) GetData() int {
	return sll.data
}

func (sll *SinglyLL) SetData(data int) {
	sll.data = data
}

func (sll *SinglyLL) GetNext() *SinglyLL {
	return sll.next
}

func (sll *SinglyLL) SetNext(next *SinglyLL) {
	sll.next = next
}

func TraverseSinglyLinkedList(head *SinglyLL) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func GetLength(head *SinglyLL) int {
	temp := head
	length := 0
	for temp != nil {
		temp = temp.GetNext()
		length = length + 1
	}
	return length
}

func InsertAtBeginning(head **SinglyLL, data int) {
	newNode := &SinglyLL{
		data: data,
		next: *head,
	}
	*head = newNode
}
func InsertAtPosition(head **SinglyLL, data int, position int) error {
	if position < 0 {
		return errors.New("invalid position")
	}

	newNode := &SinglyLL{data: data}
	if position == 0 {
		newNode.SetNext(*head)
		*head = newNode
		return nil
	}

	current := *head
	for i := 0; i < position-1; i++ {
		if current.GetNext() == nil {
			return errors.New("position out of range")
		}
		current = current.next
	}

	newNode.SetNext(current.GetNext())
	current.SetNext(newNode)
	return nil
}

func InsertAtEnd(head **SinglyLL, data int) error {
	newNode := &SinglyLL{data: data, next: nil}
	if *head == nil {
		*head = newNode
		return nil
	}

	current := *head
	for current.GetNext() != nil {
		current = current.GetNext()
	}
	current.SetNext(newNode)
	return nil
}

func DeleteFirstNode(head **SinglyLL) error {
	if *head == nil {
		return errors.New("nothing to delete")
	}
	nextNode := (*head).GetNext()
	(*head).SetNext(nil)
	*head = nextNode
	return nil
}

func DeleteLastNode(head **SinglyLL) error {
	if *head == nil {
		return errors.New("nothing to delete")
	}

	current := *head
	var previous *SinglyLL

	for current.GetNext() != nil {
		previous = current
		current = current.GetNext()
	}

	if previous == nil {
		*head = nil
		return nil
	}

	previous.SetNext(nil)
	return nil
}

func DeleteIntermediateNode(head **SinglyLL, value int) error {
	if *head == nil {
		return errors.New("nothing to delete")
	}
	if (*head).GetData() == value {
		*head = (*head).GetNext()
		return nil
	}
	current := *head
	var previous *SinglyLL
	for current != nil {
		if current.GetData() == value {
			previous.SetNext(current.GetNext())
			return nil
		}
		previous = current
		current = current.GetNext()
	}
	return errors.New("not found")
}
