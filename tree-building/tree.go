package tree

import (
	"sort"
	"errors"
)

type Node struct {
	ID int
	Children []*Node
}

type Record struct {
	ID int
	Parent int
}

func (node *Node) FindChildren(records []Record) (out []*Node) {

	var children []*Node
	for _, record := range records {
		if node.ID == record.Parent && node.ID != record.ID {
			node := Node{ID:record.ID}
			children = append(children, &node)
		}
	}

	return children
}

func (node *Node) Traverse(records []Record) {

	children := node.FindChildren(records)
	if len(children) > 0 {
		node.Children = children
	}

	for _, child := range children {
		child.Traverse(records)
	}
}

func Validate(records []Record) (out bool, err string) {

	// Sort records by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID <= records[j].ID
	})

	prev_record := records[0]
	for _, record := range records[1:] {

		if record.ID == prev_record.ID {
			return false, "duplicate records"
		}
		if record.ID != (prev_record.ID + 1) {
			return false, "non continuos"
		}
		if record.ID <= record.Parent {
			return false, "detected cycle"
		}

		prev_record = record
	}
	return true, ""
}

func CreateRoot(record Record) (out *Node) {

	var root *Node
	if record.ID == 0 && record.Parent == 0 {
		root = &Node {
			ID: record.ID,
		}
	}
	return root
}

func Build(records []Record) (out *Node, err error) {

	if len(records) <= 0 {

		return nil, nil
	}

	valid, error_reason := Validate(records)
	if !valid {

		return nil, errors.New(error_reason)
	}

	root := CreateRoot(records[0])

	if root == nil {

		return root, errors.New("invalid records")
	}

	root.Traverse(records)

	return root, nil

}