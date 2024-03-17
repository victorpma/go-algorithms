package binarytree_test

import (
	"testing"

	binarytree "github.com/victorpma/go-algorithms/binaryTree"
)

func TestBinaryTree_Search(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{"Test 1", []int{5, 3, 7, 1}},
		{"Test 2", []int{10, 5, 15, 20, 50, 2}},
		{"Test 3", []int{2, 1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &binarytree.BinaryTree{}
			for _, v := range tt.values {
				tree.Insert(v)
			}
			for _, v := range tt.values {
				target := tree.Search(v)
				if target == nil || target.Value() != v {
					t.Errorf("error on scenario %s got %v, want %v", tt.name, nil, v)
				}
			}
		})
	}
}
