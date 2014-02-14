package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var rows, cols, start, end int
var cpuProfile string

func init() {
	flag.IntVar(&rows, "r", 6, "Rows")
	flag.IntVar(&cols, "c", 6, "Cols")
	flag.IntVar(&start, "s", 0, "Start index")
	flag.IntVar(&end, "e", 34, "End index")
	flag.StringVar(&cpuProfile, "prof", "", "CPU profile output location")
	flag.Parse()
}

type Problem struct {
	Graph  *GridGraph
	Start  int
	End    int
	Solves int
}

func findHamiltonianPaths(p *Problem, s *MoveStack, path *Path) {
	var start, next int
    var adjList []int

	for {
		if s.count == 0 {
			return // p.Solves already holds amount of solves
		}
		start, next = s.Pop()

		// Using -1 to signal a pop off of the path stack
		if start == -1 {
			path.Pop()
			continue
		}

		if p.Graph.CanMove(start, next) {
			// Push to path
			path.Push(next)

			// Add pop off of the path stack
			s.Push(-1, 0)

			if next == p.End {
				if isHamiltonianPath(p.Start, p.End, p.Graph, path) {
					p.Solves++
				}
			} else {
				adjList = p.Graph.GetAdjList(next)
				for _, idx := range adjList {
					// don't add if already in path
					if path.set[idx] == 1 {
						continue
					} else {
						s.Push(next, idx)
					}
				}
			}
		}
	}
}

func main() {
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	graph := NewGridGraph(rows, cols)
	path := NewPath(graph.Num)
	path.Push(start)

	s := NewMoveStack(1000)
	adjList := graph.GetAdjList(start)
	for _, idx := range adjList {
		s.Push(start, idx)
	}

	p := &Problem{Graph: graph, Start: start, End: end}
	findHamiltonianPaths(p, s, path) // result gets stored in Problem
	fmt.Printf("# solves: %d\n", p.Solves)
}
