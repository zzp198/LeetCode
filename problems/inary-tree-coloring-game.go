package problems

func btreeGameWinningMove(root *TreeNode, n, x int) bool {
	lsz, rsz := 0, 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ls := dfs(node.Left)
		rs := dfs(node.Right)
		if node.Val == x {
			lsz, rsz = ls, rs
		}
		return ls + rs + 1
	}
	dfs(root)
	return max(max(lsz, rsz), n-1-lsz-rsz)*2 > n
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
