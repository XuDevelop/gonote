package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历  先输出root结点，在输出左子树，在输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历  先输出root的左子树，再输出root结点，在输出root的右子树
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	//构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "小明",
	}

	left1 := &Hero{
		No:   2,
		Name: "小红",
	}

	right1 := &Hero{
		No:   3,
		Name: "小蓝",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "小黄",
	}

	right1.Right = right2

	// PreOrder(root)
	InfixOrder(root)
}
