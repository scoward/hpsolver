package main

import ()

type movePair struct {
	start int
	next  int
}

type MoveStack struct {
	nodes []movePair
	count int
}

func NewMoveStack(capacity int) *MoveStack {
	return &MoveStack{nodes: make([]movePair, capacity), count: 0}
}

func (s *MoveStack) Push(start, next int) {
    if s.count >= len(s.nodes) {
		nodes := make([]movePair, len(s.nodes)*2)
		copy(nodes, s.nodes)
		s.nodes = nodes
	}
	s.nodes[s.count].start = start
	s.nodes[s.count].next = next
	s.count++
}

func (s *MoveStack) Pop() (start, next int) {
	if s.count == 0 {
		return -1, -1
	}
    s.count--
    return s.nodes[s.count].start, s.nodes[s.count].next
}
