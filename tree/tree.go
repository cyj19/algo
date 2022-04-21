/**
 * @Author: cyj19
 * @Date: 2022/4/20 16:08
 */

package tree

import (
	"fmt"
	"github.com/cyj19/algo/queue"
)

/**
树Tree的定义：
1、有节点间的层次关系，分为父节点和子节点
2、有唯一一个根节点，该根节点没有父节点
3、除了根节点，其他节点有且只有一个父节点
4、每个节点本身以及后代都是一个树，是一个递归的结构
5、没有后代的节点称为叶子节点，没有节点的树称为空树
*/

// TreeNode 树
type TreeNode struct {
	Data  interface{} // 节点用来存放数据
	Left  *TreeNode   // 左子树
	Right *TreeNode   // 右子树
}

// PreOrder 先序遍历（先访问根节点，再访问左子树，最后访问右子树）传入的参数视为根节点
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Println(tree.Data)

	// 递归遍历左子树
	PreOrder(tree.Left)
	// 递归遍历右子树
	PreOrder(tree.Right)
}

// MidOrder 中序遍历（先访问左子树，再访问根节点，最后访问右子树）传入的参数视为根节点
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 遍历左子树
	MidOrder(tree.Left)

	// 打印根节点
	fmt.Println(tree.Data)

	// 遍历右子树
	MidOrder(tree.Right)
}

// PostOrder 后序遍历（先访问左子树，再访问右子树，最后访问根节点）
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 遍历左子树
	PostOrder(tree.Left)

	// 遍历右子树
	PostOrder(tree.Right)

	// 打印根节点
	fmt.Println(tree.Data)
}

// LayerOrder 层序遍历（广度遍历，要使用队列辅助），每一层从左往右遍历每个节点，传入参数视为根节点
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 新建队列
	q := new(queue.LinkQueue)
	// 根节点入队
	q.Add(tree)
	for q.Size() > 0 {
		// 取出队列的元素并打印值
		node := (q.Remove()).(*TreeNode)

		fmt.Println(node.Data)

		// 每层节点从左往后入队
		// 左子树非空，入队
		if node.Left != nil {
			q.Add(node.Left)
		}

		// 右子树非空，入队
		if node.Right != nil {
			q.Add(node.Right)
		}
	}
}
