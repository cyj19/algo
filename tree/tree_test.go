/**
 * @Author: cyj19
 * @Date: 2022/4/20 16:23
 */

package tree

import "testing"

func TestTreeNode(t *testing.T) {
	tree := &TreeNode{Data: 0}
	tree.Left = &TreeNode{Data: 1}
	tree.Right = &TreeNode{Data: 2}

	tree.Left.Left = &TreeNode{Data: 3}
	tree.Left.Right = &TreeNode{Data: 4}

	tree.Right.Left = &TreeNode{Data: 5}
	tree.Right.Right = &TreeNode{Data: 6}

	t.Log("先序遍历")
	PreOrder(tree)

	t.Log("中序遍历")
	MidOrder(tree)

	t.Log("后序遍历")
	PostOrder(tree)

	t.Log("层序遍历")
	LayerOrder(tree)
}
