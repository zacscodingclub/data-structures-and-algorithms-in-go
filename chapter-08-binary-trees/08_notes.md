# Binary Trees

## Y Tho?

Combies advantages of two other structures: an ordered array and a linked list. You can search a tree quickly, as you can an ordered array, and you can also insert and delete items quickly, as you can with a linked list.

### What is a tree*? (* binary tree)

A tree consists of nodes connected by edges. A tree is a more general category called a graph. In programs, nodes often represent such entities as objects, people, car parts, etc. The edges represent the way the nodes are related. Loosely speaking, it's easy for a program to visit other nodes when there is an edge or reference/pointer to the node. Generally you are restricted to visiting the nodes from the root downward.

## Tree Terminology

1. Path: Think of someone walking from node to node along the edges that connect them. The resulting sequence of nodes is called a path.
2. Root: The node at the top of the tree. There is only one root in a tree.
3. Parent: Any node (except the root) has exactly one edge running upward to another node. The node above it is called the parent of the node.
4. Child: Any node may have at 1+ edge running downward to another node. The node below it is called the child of the node.
5. Leaf: A node that has no children is called a leaf node or simply a leaf. There can only be one root in a tree but there can be many leaves.
6. Subtree: Any node may be considered to be the root of a subtree, which consists of its children and its children's children, and so on. If you think in terms of families, a node's subtree contains all its descendants.
7. Visiting: A node is visited when a program arrives at the node and performs an operation, e.g. checking value of attribute or displaying. Merely passing over a node is not considered visiting.
8. Traversing: To traverse a tree means to visit all the nodes in a specified order. For example, you might visit all the nodes in order of ascending key value.
9. Levels: The level of a particular node refers to how many generations the node is from the root. Root is Level 0, it's children Level 1, and so forth.
10. Keys: Used to search for item or perform other operations on it.
11. Binary Trees: If every node in a tree can have at most two children, the tree is called a binary tree. The two children are called the left and right child. A node could have 2 children, a left or right child, or no children (it's a leaf).

## An Analogy

Heirarchical file structure in a computer system. It does differ, but this is a good mental model to approach it.

## How do Binary Search Trees Work?

### Unbalanced Tree

An unbalanced tree is when one "side" of the root has more total nodes than the other. Trees can become unbalanced because of insertion order.

## Finding a Node

Loops through starting at the root to determine if a key is equal to the value stored in the node. If the current value is less than the key, it will traverse the left edge and vice versa.

## Inserting a Node

Find the place to insert then insert into place by correcting the connections.

## Traversing the Tree

Visiting each node in a specified order. Not particularly fast but interesting in it's own right.

### Inorder

ascending order based on their key values, helpful for creating a sorted list of data. Simplest way is via recursion.

### Preorder & Postorder

Often used for analyzing algabraic expressions. Revisit prefix/postfix conversions into queues/stacks.

## Finding Maximum and Minimum Values

Very effecient at determining this value.

## Deleting a Node

Most complicated common operation requried for binary search trees.

## Efficiency

## Trees Represented as Arrays

## Duplicate Keys

## Complete Tree

## The Huffman Code
