package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RopeNode struct {
	Left   *RopeNode
	Right  *RopeNode
	Weight int
	Data   string
}

type QueueNode struct {
	next *QueueNode
	val  *RopeNode
}

type Queue struct {
	len  int
	head *QueueNode
	tail *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Front() *RopeNode {
	if q.len > 0 {
		nodeVal := q.head.val
		q.head = q.head.next
		q.len--
		return nodeVal
	}
	return nil
}

func (q *Queue) Push(r *RopeNode) {
	node := &QueueNode{val: r}
	if q.len > 0 {
		q.tail.next = node
	} else {
		q.head = node
	}
	q.tail = node
	q.len++
}

func NewRopeBuilder() *RopeNode {
	return &RopeNode{}
}

func chunkizeInput(s string) []string {
	CHUNK_SIZE := 2
	var chunks []string
	for i := 0; i < len(s)-1; i += CHUNK_SIZE {
		fmt.Println(i)
		chunks = append(chunks, s[i:min(len(s)-1, i+CHUNK_SIZE)])
	}
	fmt.Println("No. of chunks: ", len(chunks))
	return chunks
}

func isLeaf(r *RopeNode) bool {
	if r.Left == nil && r.Right == nil {
		return true
	}
	return false
}

func insertChunks(r *RopeNode, chunks []string, idx *int) int {
	if r == nil {
		return 0
	}
	leftWeight := insertChunks(r.Left, chunks, idx)
	rightWeight := insertChunks(r.Right, chunks, idx)
	if isLeaf(r) {
		r.Weight = len(chunks[*idx])
		r.Data = chunks[*idx]
		*idx++
		return r.Weight
	}
	r.Weight = leftWeight
	return leftWeight + rightWeight
}

func levelOrder(r *RopeNode) {
	if r == nil {
		return
	}
	q := NewQueue()
	q.Push(r)
	for q.len > 0 {
		qLen := q.len
		fmt.Println("Queue len ", q.len)
		for qLen > 0 {
			node := q.Front()
			fmt.Printf("%d %s", node.Weight, node.Data)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			qLen--
		}
	}
}

func main() {
	rope := NewRopeBuilder()
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	chunks := chunkizeInput(line)
	fmt.Println(chunks)
	root := rope
	rope.Left = NewRopeBuilder()
	rope.Left.Left, rope.Left.Right = NewRopeBuilder(), NewRopeBuilder()
	rope.Left.Left.Left, rope.Left.Left.Right = NewRopeBuilder(), NewRopeBuilder()
	rope.Left.Right.Left, rope.Left.Right.Right = NewRopeBuilder(), NewRopeBuilder()
	idx := 0
	l1 := insertChunks(root, chunks, &idx)
	fmt.Println(l1)
	levelOrder(root)
}
