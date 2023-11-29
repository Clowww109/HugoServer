package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

type Node2 struct {
	Key    int
	Height int
	Left   *Node2
	Right  *Node2
}

type AVLTree struct {
	Root *Node2
}

func NewNode(key int) *Node2 {
	return &Node2{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func height(node *Node2) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node2) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func getBalance(node *Node2) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node2) *Node2 {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func rightRotate(y *Node2) *Node2 {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func insert(node *Node2, key int) *Node2 {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func GenerateTree(count int) *AVLTree {
	avl := AVLTree{}
	for i := 0; i <= count; i++ {
		key := rand.Intn(100)
		avl.Insert(key)
	}
	return &avl
}

func GetMermaidInLines(node *Node2, result *string) {
	if node != nil {
		if node.Left != nil {
			*result += fmt.Sprintf("%d --> %d\n", node.Key, node.Left.Key)
		}
		if node.Right != nil {
			*result += fmt.Sprintf("%d --> %d\n", node.Key, node.Right.Key)
		}
		GetMermaidInLines(node.Left, result)
		GetMermaidInLines(node.Right, result)
	}
}

func GetMermaidAVL(count int) string {
	tree := GenerateTree(count)
	mermaidSummary := ""
	GetMermaidInLines(tree.Root, &mermaidSummary)
	return mermaidSummary
}

func GetTask2Page(treeCount int) string {
	filePath := "/app/static/tasks/binary.md"
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	oldData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	str := strings.Split(string(oldData), "\n")
	for i, s := range str {
		if strings.Contains(s, "Как только дерево достигнет 100 элементов") {
			str = str[:i+1]
			break
		}
	}
	var graphStr = ""
	graphStr = fmt.Sprintf("\n{{< mermaid >}}\ngraph TD\n %s {{< /mermaid >}}", GetMermaidAVL(treeCount))
	str = append(str, graphStr)

	newData := strings.Join(str, "\n")

	return newData
}
