<?php
class Node
{
    public string $data;
    public ?Node $next = null;

    public function __construct($data)
    {
        $this->data = $data;
    }
}

class LinkedList
{
    private Node $head;
    private int $length = 0;

    public function insert(Node $newNode)
    {
        if ($this->length === 0) {
            $this->head = $newNode;
            $this->length++;
            return;
        }

        $currentNode = $this->head;
        while (!empty($currentNode->next)) {
            $currentNode = $currentNode->next;
        }

        $currentNode->next = $newNode;
        $this->length++;
    }

    public function remove(string $query)
    {
        if ($this->length < 1) {
            return;
        }

        $previousNode = null;
        $currentNode = $this->head;

        for ($i = 0; $i < $this->length; $i++) {
            $previousNode = $currentNode;

            if ($currentNode->data === $query) {
                $previousNode->next = $currentNode->next;
                $this->length--;
                unset($currentNode);

                return;
            }

            $currentNode = $currentNode->next;
        }

        echo ("$query not found in list.\n");
    }

    public function display()
    {
        $currentNode = $this->head;
        for ($i = 0; $i < $this->length; $i++) {
            echo ($currentNode->data . "\n");
            $currentNode = $currentNode->next;
        }

        echo ("\n");
    }
}

function main()
{
    $list = new LinkedList();

    $node1 = new Node("Node One");
    $list->insert($node1);

    $node2 = new Node("Node Two");
    $list->insert($node2);

    $node2 = new Node("Node Three");
    $list->insert($node2);
    $list->display();

    $list->remove("Node Two");
    $list->display();

    $list->remove("Node five");
}

main();
