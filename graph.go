package main

import ()

type GridGraph struct {
	Num       int
	Rows      int
	Cols      int
	AdjMatrix [][]int
	AdjShort  [][]int // holds a list of adjacent indices
}

func NewGridGraph(rows, cols int) *GridGraph {
	num := rows * cols
	adj := make([][]int, num)
	adjShort := make([][]int, num)

	for idx, _ := range adj {
		adj[idx] = make([]int, num)
		adjShort[idx] = make([]int, 2) // min length possible is 2
		col := idx % cols
		row := idx / cols

		// Check if the vertex can move up, down, left, or right based on index.
		// If possible, mark true (1) in the adjacency matrix.
		if col != 0 {
			adj[idx][idx-1] = 1 // can move to a left vertex
			adjShort[idx] = append(adjShort[idx], idx-1)
		}
		if col < cols-1 {
			adj[idx][idx+1] = 1 // can move to a right vertex
			adjShort[idx] = append(adjShort[idx], idx+1)
		}
		if row != 0 {
			adj[idx][idx-cols] = 1 // can move up a vertex
			adjShort[idx] = append(adjShort[idx], idx-cols)
		}
		if row < rows-1 {
			adj[idx][idx+cols] = 1 // can move down a vertex
			adjShort[idx] = append(adjShort[idx], idx+cols)
		}
	}

	graph := GridGraph{Rows: rows, Cols: cols, Num: num, AdjShort: adjShort, AdjMatrix: adj}
	return &graph
}

func (g *GridGraph) GetAdjList(idx int) []int {
	return g.AdjShort[idx]
}

func (g *GridGraph) CanMove(from, to int) bool {
	return g.AdjMatrix[from][to] == 1
}
