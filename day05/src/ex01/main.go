package main

import (
	"day05/tree"
	"fmt"
)

func example() *tree.TreeNode {
	t := &tree.TreeNode{
		HasToy: true,
		Left: &tree.TreeNode{
			HasToy: true,
			Left: &tree.TreeNode{
				HasToy: true,
			},
			Right: &tree.TreeNode{
				HasToy: false,
			},
		},

		Right: &tree.TreeNode{
			HasToy: false,
			Left: &tree.TreeNode{
				HasToy: true,
			},
			Right: &tree.TreeNode{
				HasToy: true,
			},
		},
	}

	return t
}

func test() *tree.TreeNode {
	t := &tree.TreeNode{
		HasToy: true,

		Left: &tree.TreeNode{
			HasToy: false,
			Left: &tree.TreeNode{
				HasToy: false,
			},
			Right: &tree.TreeNode{
				HasToy: false,
			},
		},

		Right: &tree.TreeNode{
			HasToy: false,
			Right: &tree.TreeNode{
				HasToy: true,
			},
			Left: &tree.TreeNode{
				HasToy: true,
			},
		},
	}

	return t
}

func unrollGarland(root *tree.TreeNode) []bool {
	if root == nil {
		return nil
	}

	ans := make([]bool, 0)
	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)
	lvl := 1
	for len(queue) > 0 {
		size := len(queue)
		toys := make([]bool, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			toys = append(toys, node.HasToy)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		if lvl%2 == 1 {
			for i, j := 0, len(toys)-1; i < j; i, j = i+1, j-1 {
				toys[i], toys[j] = toys[j], toys[i]
			}
		}

		ans = append(ans, toys...)
		lvl++
	}

	return ans
}

func main() {
	example, test := unrollGarland(example()), unrollGarland(test())
	fmt.Printf("example from readme: %t\nmy test: %t\n", example, test)
}
