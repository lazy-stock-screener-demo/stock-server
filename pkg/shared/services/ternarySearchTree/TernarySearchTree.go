package ternarySearchTree

import (
	"fmt"
	"strings"
)

type Node struct {
	value  string
	r      rune
	endTag bool
	left   *Node
	right  *Node
	middle *Node
}

type TreeProps struct {
	WordList []string
}

type Tree struct {
	node      *Node
	result    []string
	resultMap map[string]bool
}

func (a *Tree) Init(wordList []string) {
	for _, word := range wordList {
		a.node = a.insert(word, a.node)
	}
}

func (a *Tree) insert(word string, node *Node) *Node {
	char := string(word[0])
	r := rune(word[0])
	if node.value == "" {
		node.value = char
		node.r = r
	}
	if r < node.r {
		if node.left == nil {
			node.left = &Node{
				value:  "",
				endTag: false,
			}
		}
		node.left = a.insert(word, node.left)
	} else if r > node.r {
		if node.right == nil {
			node.right = &Node{
				value:  "",
				endTag: false,
			}
		}
		node.right = a.insert(word, node.right)
	} else {
		if len(word) == 1 {
			node.endTag = true
			return node
		}
		if node.middle == nil {
			node.middle = &Node{
				value:  "",
				endTag: false,
			}
		}
		node.middle = a.insert(word[1:], node.middle)
	}
	return node
}

func (a *Tree) findAllSuffixes(pattern string, node *Node, results []string) {
	if node == nil {
		return
	}
	if node.endTag {
		str := fmt.Sprintf("%s%s", pattern, node.value)
		// fmt.Println("matching", str)
		a.result = append(a.result, str)
		a.resultMap[str] = true
	}
	a.findAllSuffixes(pattern, node.left, results)
	a.findAllSuffixes(pattern, node.right, results)
	a.findAllSuffixes(fmt.Sprintf("%s%s", pattern, node.value), node.middle, results)
}

var resultList []string

func (a *Tree) Search(chars string) []string {
	var result []string
	charList := strings.Split(chars, "")
	for _, char := range charList {
		resultList = append(result, a.findNode(char)...)
	}
	for word := range a.resultMap {
		result = append(result, word)
	}
	// fmt.Println("resultList", resultList)
	a.result = []string{}
	a.resultMap = map[string]bool{}
	return result
}

func (a *Tree) findNode(pat string) []string {
	node := a.node
	for _, runedChar := range pat {
		for true {
			if runedChar > node.r {
				node = node.right
			} else if runedChar < node.r {
				node = node.left
			} else {
				node = node.middle
				break
			}
			if node == nil {
				break
			}
		}
	}
	a.findAllSuffixes(pat, node, []string{})
	return a.result
}

func NewTree() *Tree {
	return &Tree{
		node: &Node{
			value:  "",
			endTag: false,
		},
		resultMap: map[string]bool{},
	}
}
