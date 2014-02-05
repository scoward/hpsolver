package main

import ()

type GridGraph struct {
	Num  int
	Rows int
	Cols int
}

func (g *GridGraph) CanMove(from, to int) bool {
	fromRow := from / g.Cols
	fromCol := from % g.Cols
	toRow := to / g.Cols
	toCol := to % g.Cols

	diffCol := toCol - fromCol
	if diffCol < 0 {
		diffCol = -diffCol
	}
	diffRow := toRow - fromRow
	if diffRow < 0 {
		diffRow = -diffRow
	}

	diff := diffCol + diffRow
	// If diff is > 1 then there is at least one vertex between from and to
	if diff != 1 {
		return false
	}

	return true
}

func NewGridGraph(rows, cols int) *GridGraph {
	num := rows * cols
	graph := GridGraph{Rows: rows, Cols: cols, Num: num}
	return &graph
}
