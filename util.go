package main

import (
)

func isHamiltonianPath(start, end int, graph *GridGraph, path *Path) bool {
    if path.Count() != graph.Num {
        return false
    }
    if start != path.Bottom() {
        return false
    }
    if end != path.Peek() {
        return false
    }

    return true
}
