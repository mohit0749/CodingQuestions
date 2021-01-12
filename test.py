# Python3 program to flip
# a binary tree

# A binary tree node
class Node:

    # Constructor to create
    # a new node
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.right = right
        self.left = left


def flipBinaryTree(root):
    # Base Cases
    if root is None:
        return root

    if (root.left is None or
            root.right is None):
        return root

    # Recursively call the
    # same method
    flippedRoot = flipBinaryTree(root.left)

    # Rearranging main root Node
    # after returning from
    # recursive call
    root.left.left = root.right
    root.left.right = root
    root.left = root.right = None

    return flippedRoot


# Iterative method to do the level
# order traversal line by line
def printLevelOrder(root):
    # Base Case
    if root is None:
        return

    # Create an empty queue for
    # level order traversal
    from queue import Queue
    q = Queue()

    # Enqueue root and initialize
    # height
    q.put(root)

    while (True):

        # nodeCount (queue size) indicates
        # number of nodes at current level
        nodeCount = q.qsize()
        if nodeCount == 0:
            break

        # Dequeue all nodes of current
        # level and Enqueue all nodes
        # of next level
        while nodeCount > 0:
            node = q.get()
            print(node.data, end=' ')
            if node.left is not None:
                q.put(node.left)
            if node.right is not None:
                q.put(node.right)
            nodeCount -= 1
        print()


# Driver code
root = Node(1)
root.left = Node(2, Node(3, Node(4, None, Node(5))), Node(6))
root.right = Node(9, Node(8), Node(7))

print("Level order traversal of given tree")
printLevelOrder(root)

root = flipBinaryTree(root)

print("\nLevel order traversal of the flipped tree")
printLevelOrder(root)

# This code is contributed by Nikhil Kumar Singh(nickzuck_007)

if __name__ == '__main__':
    pass
