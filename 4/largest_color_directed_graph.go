package main

import (
	"fmt"
	"sort"
)

func main() {
	count := LargestPathValue("hhqhuqhqff", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}})

	fmt.Printf("maxSize: %v\n", count)
}

func LargestPathValue(colors string, edges [][]int) int {
	var graph Graph

	graph.buildGraph(colors, edges)

	longest := graph.longestPath()

	return graph.mostRepeatedColorCount(longest)
}

type Node struct {
	index        int
	color        string
	OutNeighbors []*Node
	InNeighbors  []*Node
}

type Graph struct {
	nodes    map[int]*Node
	hasCycle bool
}

func (g *Graph) addNode(nodeID int, color string) *Node {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := &Node{color: color}
		g.nodes[nodeID] = newNode
		return newNode
	} else {
		return g.nodes[nodeID]
	}
}

func (g *Graph) buildGraph(colors string, edges [][]int) {
	g.nodes = make(map[int]*Node)
	for _, ed := range edges {
		begin := g.addNode(ed[0], string(colors[ed[0]]))
		end := g.addNode(ed[1], string(colors[ed[1]]))

		begin.OutNeighbors = append(begin.OutNeighbors, end)
		end.InNeighbors = append(end.InNeighbors, begin)
	}
}

func (g *Graph) longestPath() []*Node {
	visited := make(map[*Node]bool)
	recStack := make(map[*Node]bool)
	return g.dfs(g.nodes[0], visited, recStack)
}

func (g *Graph) dfs(node *Node, visited map[*Node]bool, recStack map[*Node]bool) []*Node {
	if visited[node] {
		return nil
	}

	visited[node] = true
	recStack[node] = true

	var maxPath []*Node

	if node.OutNeighbors != nil {
		for _, neighbor := range node.OutNeighbors {
			if recStack[neighbor] {
				g.hasCycle = true
				return nil
			}

			path := append(g.dfs(neighbor, visited, recStack), node)
			if len(path) > len(maxPath) {
				maxPath = path
			}
		}

		recStack[node] = false
		return maxPath
	}

	return []*Node{node}
}

func (g *Graph) mostRepeatedColorCount(path []*Node) int {
	if g.hasCycle {
		return -1
	}

	colorCountMap := make(map[string]int)
	for _, node := range path {
		_, exists := colorCountMap[node.color]

		if exists {
			colorCountMap[node.color] += 1
		} else {
			colorCountMap[node.color] = 1
		}
	}

	var list []int
	for _, val := range colorCountMap {
		list = append(list, val)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	return list[0]
}
