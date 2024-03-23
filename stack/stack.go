package stack 


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

