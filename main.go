package main

import (
	"fmt"
)

type RopeNode struct {
  Left *RopeNode
  Right *RopeNode
  Weight int 
  Data string 
  TotalWeight int 
}

type stackNode struct {
  next *stackNode
  val *RopeNode
}

type Stack struct {
  len int
  head *stackNode 
}

func NewStack() *Stack {
  return &Stack{}
}

func (s *Stack) Len() int {
  return s.len
}

func (s *Stack) Peek() *RopeNode {
  if s.Len() > 0 {
    return s.head.val
  }
  return nil
}

func (s *Stack) Push(val *RopeNode) {
  node := &stackNode{val: val}
  node.next = s.head 
  s.head = node
  s.len++
}

func (s *Stack) Pop() *RopeNode {
  if s.Len() > 0 {
    node := s.head 
    s.head = s.head.next
    s.len--
    return node.val
  }
  return nil
}


func NewRopeBuilder() *RopeNode {
  return &RopeNode{}
}


func getString(r *RopeNode) string {
  if r == nil {
    return ""
  }
  leftSubstring := getString(r.Left)
  leftSubstring = fmt.Sprintf("%s%s", leftSubstring, r.Data)
  fmt.Println(r.Weight, r.Data)
  rightSubstring := getString(r.Right)
  return fmt.Sprintf("%s%s", leftSubstring, rightSubstring)
}

func BuildTree(r *RopeNode) int {
  if r == nil {
    return 0 
  }
  leftWeight := BuildTree(r.Left)
  rightWeight := BuildTree(r.Right)
  r.TotalWeight = leftWeight + rightWeight + len(r.Data)
  r.Weight = leftWeight + len(r.Data)
  return r.TotalWeight
}

func isLeaf(r *RopeNode) bool {
  return r.Left == nil && r.Right == nil
}

func (r *RopeNode) insert (idx int, s string){
  if r == nil {
    return 
  }
  if isLeaf(r) {
    if len(r.Data) >= idx {
      before := r.Data[:idx]
      after := r.Data[idx:]
      r.Data = fmt.Sprintf("%s%s%s", before, s, after)
    }
  }
  if idx <= r.Weight {
    if r.Left != nil {
      r.Left.insert(idx, s)
    }
  } else {
    if r.Right != nil {
      r.Right.insert(idx-r.Weight, s)
    }
  }

}

func main() {
  Rope := NewRopeBuilder()
  root := Rope
  Rope.Left = &RopeNode{}
  Rope = Rope.Left
  Rope.Left = &RopeNode{Data: "Hello "}
  Rope.Right = &RopeNode{Data: "World!"}
  BuildTree(root)
  root.insert(6, "Great")
  BuildTree(root)
  getString(root)
}
