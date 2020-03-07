package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID int
	Children []*Node
}

func (n *Node) AddChild(c *Node) {
	n.Children = append(n.Children, c)
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].Parent != 0 || records[0].ID != 0 {
		return nil, fmt.Errorf("mal-formatted root node")
	}

	node := Node{
		ID:       0,
		Children: nil,
	}

	nodes := make(map[int]*Node, len(records))
	nodes[0] = &node

	for _, r := range records[1:] {
		outOfBoundsId := r.ID < 0 || r.ID >= len(records)
		invalidIdNumber := r.ID <= r.Parent || r.ID == 0
		duplicatedId := nodes[r.ID] != nil
		if outOfBoundsId || invalidIdNumber || duplicatedId {
			return nil, fmt.Errorf("invalid record { ID: %d, Parent: %d }", r.ID, r.Parent)
		}

		curNode := Node{
			ID:       r.ID,
			Children: nil,
		}
		nodes[r.ID] = &curNode
		nodes[r.Parent].AddChild(&curNode)
	}
	return &node, nil
}
