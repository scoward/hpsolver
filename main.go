package main

import (
	"fmt"
)

type Problem struct {
	Graph  *GridGraph
	Start  int
	End    int
	Solves int
}

func findHamiltonianPaths(p *Problem, path *Path, startIdx, endIdx int) {
	path.Push(startIdx)
	for next := 0; next < p.Graph.Num; next++ {
		if path.Contains(next) {
			continue
		}
		if !p.Graph.CanMove(startIdx, next) {
			continue
		}

		if next == endIdx {
			path.Push(next)
			if isHamiltonianPath(p.Start, p.End, p.Graph, path) {
				p.Solves++
			}
			path.Pop()
		} else {
			findHamiltonianPaths(p, path, next, endIdx)
		}
	}
	path.Pop()
}

func main() {
	numRows := 7
	numCols := 7

	graph := NewGridGraph(numRows, numCols)
	p := &Problem{Graph: graph, Start: 0, End: 48}
	path := NewPath(graph.Num)
	findHamiltonianPaths(p, path, p.Start, p.End) // result gets stored in Problem
	fmt.Printf("# solves: %d\n", p.Solves)
}
