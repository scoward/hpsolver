package main

import (
	"testing"
    "fmt"
)

func Test6x6(t *testing.T) {
    graph := NewGridGraph(6, 6)
    p := &Problem{Graph: graph, Start: 0, End: 34}
    path := NewPath(graph.Num)
    findHamiltonianPaths(p, path, p.Start, p.End) // result gets stored in Problem

    if p.Solves != 590 {
        t.Error("For 6x6 expected 590 got", p.Solves)
    }
}

func Benchmark6x6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		graph := NewGridGraph(6, 6)
		p := &Problem{Graph: graph, Start: 0, End: 34}
		path := NewPath(graph.Num)
		findHamiltonianPaths(p, path, p.Start, p.End) // result gets stored in Problem

		if p.Solves != 590 {
			fmt.Printf("ERROR: 6x6 benchmark solves not 590, instead %d\n", p.Solves)
		}
	}
}
