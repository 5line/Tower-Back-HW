package main

import "fmt"

type Node struct{
	Data int
	Left *Node
	Right *Node
}


type BST struct{
	Root *Node
}

func (bst *BST) Insert(data int) {
	newNode := &Node{Data: data}
	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	current := bst.Root
	for {
		if data < current.Data {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
		current = current.Right	
		}
	}
}


func (bst *BST) Delete(data int) {
	if bst.Root == nil {
		return
	}
	current := bst.Root
	parent := &Node{}

	for {
		if data < current.Data {
			parent = current.Left
		} else if data > current.Data {
			parent = current
			current = current.Right
		} else {
			break
		}
		if current == nil{
			return
		}
	}

	if current.Left == nil && current.Right == nil {
		if parent.Data > current.Data {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if current.Left == nil { //узел  имеет только правого ребенка
		if parent.Data > current.Data {
			parent.Left = current.Right
		} else {
			parent.Right = current.Right
		}
		} else if current.Right == nil { // Узел имеет только левого ребенка
			if parent.Data > current.Data {
			 parent.Left = current.Left
			} else {
			 parent.Right = current.Left
			}
		} else { // Узел имеет двух детей
			successor := current.Right
			successorParent := current
			for successor.Left != nil {
			 successorParent = successor
			 successor = successor.Left
			}
		  
			if successorParent != current {
			 successorParent.Left = successor.Right
			} else {
			 successorParent.Right = successor.Right
			}
		  
			if parent.Data > current.Data {
			 parent.Left = successor
			} else {
			 parent.Right = successor
			}
			successor.Left = current.Left
			successor.Right = current.Right
		}
}
		  
func (bst *BST) Contains(data int) bool {
	current := bst.Root
	for current != nil {
		if data < current.Data {
			current = current.Left
		} else if data > current.Data {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

func main() {
	var command string
	var element int
	bst := BST{}
   
	fmt.Println("У вас есть уникальная возможность добавить, удалить или проверить наличие элемента в ебучем бинарном дереве, введите команду 'Add', 'Delete' или 'Contains' соответственно, чтобы совершить желаемое действие. Также Вы можете нажать ентер, когда Вам надоест это дерьмо")
   
	for {
		fmt.Scanf("%s", &command)
		if command == "" {
			break
		}
   
		fmt.Println("Введите значение:")
		fmt.Scanf("%d\n", &element)
   
		switch command {
			case "Add":
				bst.Insert(element)
				fmt.Println("Элемент добавлен.")
			case "Delete":
				bst.Delete(element)
				fmt.Println("Элемент удален.")
			case "Contains":
				if bst.Contains(element) {
					fmt.Println("Элемент существует.")
				} else {
					fmt.Println("Элемента нет.")
	  			}
	 		default:
				fmt.Println("Еблан?")
		}
	}
	fmt.Println("Выход из программы.")
}