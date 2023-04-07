```c
bool isMirror(struct TreeNode* root1, struct TreeNode* root2) {
  if (root1 == NULL && root2 == NULL) {
    return 1;
  }

  if (root1 == NULL || root2 == NULL) {
    return 0;
  }

  if (root1->val != root2->val) {
    return 0;
  }
  return isMirror(root1->left, root2->right) && isMirror(root1->right, root2->left);
}

bool isSymmetric(struct TreeNode* root) {
  if (root == NULL) {
    return 1;
  }
  return isMirror(root->left, root->right);
}

```

```go
func isMirror(root1 *TreeNode, root2 *TreeNode) bool {
    if (root1 == nil && root2 == nil) {
        return true
    }
    // one of root is nil, different
    if (root1 == nil || root2 == nil) {
        return false
    }
    if (root1.Val != root2.Val) {
        return false
    }
    return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}


func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isMirror(root.Left, root.Right)
}

```
