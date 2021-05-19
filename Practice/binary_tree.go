package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val{
		return root
	} else if root.Val > val {
		if root.Left != nil {
			return searchBST(root.Left, val)
		}
		return nil
	} else if root.Val < val {
		if root.Right != nil {
			return searchBST(root.Right, val)
		}
		return nil
	}else{
		return nil
	}

}
//Input: root = [4,2,7,1,3], val = 2
//Output: [2,1,3]