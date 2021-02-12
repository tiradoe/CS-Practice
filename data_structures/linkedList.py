class Node:
    def __init__(self, data, ):
        self.data = data
        self.next = None


class LinkedList:
    def __init__(self, head):
        self.head = head

    def insertNode(self, newNode):
        currentNode = self.head
        while currentNode.next != None:
            currentNode = currentNode.next

        currentNode.next = newNode

    def deleteNode(self, query):
        currentNode = self.head
        previousNode = None

        while currentNode is not None:
            if currentNode.data == query:
                previousNode.next = currentNode.next
                return

            previousNode = currentNode
            currentNode = currentNode.next

        print('Could not find "%s" in queue' % query)

    def displayData(self):
        currentItem = self.head

        while currentItem is not None:
            print(currentItem.data)
            currentItem = currentItem.next


def main():
    node1 = Node('node 1')
    queue = LinkedList(node1)

    node2 = Node('node 2')
    queue.insertNode(node2)

    node3 = Node('node 3')
    queue.insertNode(node3)

    queue.displayData()

    queue.deleteNode('node 2')
    queue.displayData()

    queue.deleteNode('test')


main()
