package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"io"
	"os"
	"strings"
)

type Node struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}

func GetFigureBrackets(form string) string {
	switch form {
	case "circle":
		return "((" + form + "))"
	case "rhombus":
		return "{" + form + "}"
	default:
		return "[" + form + "]"
	}
}

func GetMermaids(nodes []*Node) string {
	var res = "\n{{< mermaid >}}\ngraph LR\n"
	for i := 0; i < len(nodes); i++ {
		for _, otherNode := range nodes[i].Links {
			res += fmt.Sprintf("%s%s --> %s%s\n",
				nodes[i].Name, GetFigureBrackets(nodes[i].Form), otherNode.Name, GetFigureBrackets(otherNode.Form))
		}
	}
	res += "{{< /mermaid >}}"
	return res
}

func NewNodes(count int) []*Node {
	var nodes []*Node
	for i := 0; i < count; i++ {
		newNode := &Node{
			ID:   i,
			Name: gofakeit.FirstName(),
		}
		nodes = append(nodes, newNode)
	}
	for i := 0; i < len(nodes); i++ {
		numbForm := gofakeit.Number(0, 5)
		switch numbForm {
		case 0:
			nodes[i].Form = "circle"
		case 1:
			nodes[i].Form = "rect"
		case 2:
			nodes[i].Form = "square"
		case 3:
			nodes[i].Form = "ellipse"
		case 4:
			nodes[i].Form = "round-rect"
		case 5:
			nodes[i].Form = "rhombus"
		}
	}
	//Случайные связи)
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if len(nodes[j].Links) < 1 {
				nodes[j].Links = append(nodes[j].Links, nodes[i])
			}
			rnd := gofakeit.Number(0, 10)
			if rnd == 2 {
				nodes[i].Links = append(nodes[i].Links, nodes[j])
			}

		}
	}
	return nodes
}

func GetTask3page() string {

	filePath := "/app/static/tasks/graph.md"
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
		if strings.Contains(s, "Все ноды графа должны быть связаны") {
			str = str[:i+1]
			break
		}
	}
	//начинается рандом
	count := gofakeit.Number(5, 30)
	nodes := NewNodes(count)
	mermGraph := GetMermaids(nodes)

	str = append(str, mermGraph)
	newData := strings.Join(str, "\n")

	return newData
}
