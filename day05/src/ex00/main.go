package main

import (
	"day05/tree"
	"fmt"
)

func areToysBalanced(node *tree.TreeNode) bool{
	status := countNode(node.Left) == countNode(node.Right)
	return status
}

func countNode(node *tree.TreeNode) int{
	count := 0
	if node != nil {

		if node.HasToy{
			count++
		}

		count += countNode(node.Left)
		count += countNode(node.Right)
	}

	return count
}

func norm() *tree.TreeNode{
	t := &tree.TreeNode{

		Left: &tree.TreeNode{
			HasToy: true,
			Left: &tree.TreeNode{
				HasToy: true,
			},
			Right: &tree.TreeNode{
				HasToy: true,
			},
		},

		Right: &tree.TreeNode{
			HasToy: false,
			Left: &tree.TreeNode{
				HasToy: false,
				Right: &tree.TreeNode{
					HasToy: true,
				},
				Left: &tree.TreeNode{
					HasToy: true,
				},
			},
			Right: &tree.TreeNode{
				HasToy: true,
			},
		},
	}

	return t
}

func trash() *tree.TreeNode{
	t := &tree.TreeNode{

		Left: &tree.TreeNode{
			HasToy: true,
		},

		Right: &tree.TreeNode{
			HasToy: true,
			Left: &tree.TreeNode{
				HasToy: false,
			},
			Right: &tree.TreeNode{
				HasToy: true,
			},
		},
	}

	return t
}


func main(){
	good, bad := areToysBalanced(norm()), areToysBalanced(trash())

	fmt.Printf("good tree:\t%t\nbad tree:\t%t\n",good, bad)
}