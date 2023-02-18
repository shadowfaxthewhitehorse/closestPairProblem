package main

import "fmt"

type Node struct {
    data int
    next *Node
}

func reverseList(head *Node) *Node {
    var prev *Node = nil
    current := head

    for current != nil {
        next := current.next
        current.next = prev
        prev = current
        current = next
    }
    return prev
}

func printList(head *Node) {
    for head != nil {
        fmt.Printf("%d -> ", head.data)
        head = head.next
    }
    fmt.Println()
}

func main() {
    head := &Node{data: 1}
    head.next = &Node{data: 2}
    head.next.next = &Node{data: 3}
    head.next.next.next = &Node{data: 4}
    head.next.next.next.next = &Node{data: 5}

    fmt.Println("Original List:")
    printList(head)

    head = reverseList(head)

    fmt.Println("Reversed List:")
    printList(head)
}
