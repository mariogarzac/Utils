package utils

import (
	"bufio"
	"log"
	"os"
    "fmt"
)

func OpenFile(filename string) (*bufio.Scanner, *os.File, error) {

    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
        return nil, file, err
    }

    sc := bufio.NewScanner(file)
    return sc, file, nil
}

func ReadWholeFile(filename string) ([]byte, error){

    var data []byte
    var err error
    data, err = os.ReadFile(filename)
    if err != nil { log.Fatal(err) }


    return data, err
}

// ----------------------------------------------------------------------------

type Stack []interface{}

// Pushes an element to the top of the stack.
func (s *Stack) Push(item interface{}){
    *s = append(*s, item)
}

// Removes last element from the stack LIFO.
func (s *Stack) Pop() (interface{}, bool){

    if len(*s) == 0{
        return 0, false
    }

    index := len(*s) - 1
    item := (*s)[index]
    *s = (*s)[:index]

    return item, true
}

// Returns top most element of the stack.
func (s *Stack) Top() (interface{}){

    if len(*s) == 0{
        return 0
    }

    index := len(*s) - 1
    item := (*s)[index]

    return item
}

// Pops the last element in the Stack
func (s *Stack) Last() (interface{}){

    if len(*s) == 0{
        return 0
    }

    item := (*s)[0]
    *s = (*s)[:0]
    return item
}

func (s *Stack) Size() int{

    return len(*s)
}


// Inserts an element in a given position.
func (s *Stack) Insert(index int, item interface{}) {
    if index < 0 || index > len(*s) {
        // Invalid index
        return
    }

    *s = append(*s, nil) // Expand the slice
    copy((*s)[index+1:], (*s)[index:]) // Shift elements to the right
    (*s)[index] = item // Insert the item at the specified index
}

// ----------------------------------------------------------------------------
type TreeNode struct {
    Name string
    Parent *TreeNode
    Value int
    IsFile bool
    Children []*TreeNode
}

type Tree struct {
    Root *TreeNode
}

func NewTreeNode(name string, isFile bool, value int) *TreeNode {
    return &TreeNode{
        Name: name,
        Value: value,
        IsFile: isFile,
        Children: make([]*TreeNode, 0),
    }
}

func NewTree(rootName string, isFile bool, value int) *Tree {
    return &Tree{
        Root: NewTreeNode(rootName, isFile, value),
    }
}

func (node *TreeNode) AddChild(child *TreeNode) {
    node.Children = append(node.Children, child)
    child.Parent = node
}

func (tree *Tree) Traverse(root *TreeNode) {
    if root != nil {
        fmt.Printf("%s %d\n",root.Name, root.Value)
    }
    for _, child := range root.Children {
        tree.Traverse(child)
    }
}


