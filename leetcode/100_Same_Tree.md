```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if (p != nil && q != nil) {
        return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else {
        return false
    }
}
```

```c
bool isSameTree(struct TreeNode* p, struct TreeNode* q) {
  if (p == NULL && q == NULL) {
    return 1;
  } else if (p != NULL && q != NULL) {
    return (p->val == q->val) && isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
  } else {
    return 0;
  }
}
```

