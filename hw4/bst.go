// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import "fmt"

type Node struct{
	key		int
	value 	int
	left 	*Node
	right 	*Node
}

//!+
type tree struct {
	root	*Node
}

func (t *tree) Insert(key int, value int) {
	var newNode *Node = &Node{key: key, value: value}
	if t.root == nil {
		t.root = newNode
	} else {
		insert(t.root, newNode)
	}
}

func insert(parentNode *Node, newNode *Node) {

	if newNode.key < parentNode.key {
		if parentNode.left == nil {
			parentNode.left = newNode
		} else {
			insert(parentNode.left, newNode)
		}
	}

	if newNode.key > parentNode.key {
		if parentNode.right == nil {
			parentNode.right = newNode
		} else {
			insert(parentNode.right, newNode)
		}
	}
}

func (t *tree) Find(key int) bool {
	return find(t.root, key)
}

func find(t *Node, key int) bool {
	if t == nil {
		fmt.Println(key, "not found :(")
		return false
	}
	if t.key == key {
		fmt.Println(key, "found!")
		return true
	}
	if key < t.key {
		return find(t.left, key)
	}
	if key > t.key {
		return find(t.right, key)
	}
	fmt.Println(key, "found!")
	return true
}

func (t *tree) IsEmpty(){
	isEmpty(t)
}

func isEmpty(t *tree) bool{
	if t.root == nil{
		fmt.Println("Tree is empty")
		return true
	}
	fmt.Println("Tree is not empty")
	return false
}

func (t *tree) InOrder() {
	inOrder(t.root)
}

func inOrder(t *Node) {
	if t == nil{
		return
	}
	
	inOrder(t.left)
	fmt.Println(t.value)
	inOrder(t.right)

}

//!-

func main() {
	var t tree
	var x tree

	t.Insert(10,10)
	t.Insert(219,219)
	t.Insert(4,4)
	t.Insert(12,12)
	t.Insert(22,22)
	t.Insert(69,69)
	t.Insert(70,70)

	t.Find(10)
	t.Find(11)

	t.IsEmpty()
	x.IsEmpty()

	t.InOrder()
}