#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node {
  int data;
  struct node* left;
  struct node* right;
};

//    3
//   / \
//  2   4
// /
//1
struct node* build_1234() {
  // build root node
  struct node* root = NULL;
  root = malloc(sizeof(struct node));

  if (root == NULL) {
    printf("malloc error!\n");
    exit(1);
  }

  root->data = 3;
  root->left = NULL;
  root->right = NULL;

  // build left node and set left node
  struct node* left = NULL;
  left = malloc(sizeof(struct node));

  if (left == NULL) {
    printf("malloc error!\n");
    exit(1);
  }

  left->data = 2;
  left->left = NULL;
  left->right = NULL;

  root->left = left;

  // build right node and set right node
  struct node* right = NULL;
  right = malloc(sizeof(struct node));

  if (right == NULL) {
    printf("malloc error!\n");
    exit(1);
  }

  right->data = 4;
  right->left = NULL;
  right->right = NULL;

  root->right = right;

  // build 4 node and set to left node left
  struct node* four = NULL;
  four = malloc(sizeof(struct node));
  if (four == NULL) {
    printf("malloc error!\n");
    exit(1);
  }

  four->data = 1;
  four->right = NULL;
  four->left = NULL;
  
  left->left = four;

  return root;
}


int size_of_node(struct node* root) {
  return (root == NULL) ? 0 : (size_of_node(root->left) + size_of_node(root->right) + 1);
}


int max_depth(struct node* root) {
  // base case 
  if (root == NULL) {
    return 0;
  }
  int max_left = max_depth(root->left);
  int max_right = max_depth(root->right);

  // find which one is more deep, add it with root node
  return max_left >= max_right ? max_left + 1 : max_right + 1;
}


// search value from a binary search tree
int look_up(struct node* root, int value) {
  if (root == NULL) {
    return 0;
  } else {
	if (root->data == value) {
      return 1;
    } else if (root->data > value) {
      return look_up(root->left, value);
    } else if (root->data < value) {
      return look_up(root->right, value);
    }
  }
}

// find min value in a binary search tree
int min_value(struct node* root) {
  struct node* current = root;
  while (current->left) {
    current = current->left;
  }
  return current->data;
}

// find max value in a binary search tree
int max_value(struct node* root) {
  struct node* current = root;
  while (current->right) {
    current = current->right;
  }
  return current->data;
}



int main(int argc, char const *argv[]) {
  
  struct node* root = build_1234();
  fprintf(stdout, "root point at\t%p\nleft point at\t%p\nright point at\t%p\n",
    root, 
    root->left, 
    root->right);

  fprintf(stdout, "size_of_node of root is %d\n", size_of_node(root));
  fprintf(stdout, "max_depth of root is %d\n", max_depth(root));
  int i;
  for (i = 0; i <= 7; i++) {
    fprintf(stdout, "find value %d res is %d\n", i, look_up(root, i));
  }


  fprintf(stdout, "find min value is %d\n", min_value(root));
  fprintf(stdout, "find max value is %d\n", max_value(root));

  free(root->left->left);
  free(root->left);
  free(root->right);
  free(root);
  return 0;
}

