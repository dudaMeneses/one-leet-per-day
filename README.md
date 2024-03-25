# One Leet per Day

> I decided to categorize the problems by algorithm implemented in the solution, thus I can add some explanation about the algorithm and the problem.

## Algorithms

### Two Pointers

> The two-pointer technique is a pattern where two pointers iterate over the data structure in tandem or separately until they satisfy a certain condition.

###### pseudocode 
```go
l, r := 0, len(nums)-1

for l < r {
    if condition == target {
    	// return some result
    } else if conditionShrinkLeft {
        // move left pointer closer to center
        l++
    } else {
    	// move right pointer closer to center
        r--
    }
}
```

### Sliding Window

> Sliding Window Technique is a method used to efficiently solve problems that involve defining a window or range in the input data (arrays or strings) and then moving that window across the data to perform some operation within the window

###### pseudocode
```go
l, r := 0, 0

for r < len(nums) {
    // expand the right part of the window
    // do something
    r++

    for condition {
        // shrink the left part of the window
        // do something
        l++
    }
}
```

### Graph Algorithms

Basically a section for algorithms that are used to traverse graphs. For that we will use the following graph representation:

#### BFS

> The breadth-first search or BFS algorithm is used to search a tree or graph data structure for a node that meets a set of criteria. 
> It begins at the root of the tree or graph and investigates all nodes at the current depth level before moving on to nodes at the next depth level.

###### pseudocode
```go
func BFS(g *Graph, startVertex *Vertex) {
    // create a queue and add the starting vertex to it
    queue := &queue{}
    queue.enqueue(startVertex)
    
    // create a map to keep track of visited vertices
    visitedVertices := map[int]bool{}
    
    // while the queue is not empty
    for queue.size() > 0 {
        // get the first vertex from the queue
        current := queue.dequeue()
        
        // visit the current vertex
        visitedVertices[current.Key] = true
        
        // for each neighboring vertex, push it to the queue if it wasn't already visited
        for _, v := range current.Vertices {
            if !visitedVertices[v.Key] {
                queue.enqueue(v)
            }
        }
    }
}
```

#### DFS

> Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. 
> The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking. Extra memory, usually a stack, is needed to keep track of the nodes discovered so far along a specified branch which helps in backtracking of the graph.

###### pseudocode
```go
func DFS(g *Graph, startVertex *Vertex) {
    // we maintain a map of visited nodes to prevent visiting the same node more than once
    visited := map[int]bool{}
    
    visited[startVertex.Key] = true
    
    // for each of the adjacent vertices, call the function recursively if it hasn't yet been visited
    for _, v := range startVertex.Vertices {
        if visited[v.Key] {
            continue
        }
        DFS(g, v)
    }
}
```

#### Dijkstra

> Dijkstra's algorithm finds the shortest path from one vertex to all other vertices. 
> It does so by repeatedly selecting the nearest unvisited vertex and calculating the distance to all the unvisited neighboring vertices.

###### pseudocode
```go
func Dijkstra(Graph, source) {
    for _, v := range Graph.Vertices{
        distance[v] = INFINITE // initial distance from source to v is infinite
        previous[v] = UNDEFINED // previous node in optimal path from source
	}
	distance[source] := 0 // distance from source to source
	Q := // the set of all nodes in Graph
	for Q is not empty {
        smallest := // node in Q with smallest distance
		remove(smallest, Q)
		for _, neighbor := range smallest.neighbors { // where neighbor has not yet been removed from Q.
            alt := dist[smallest] + length(smallest, neighbor)
            if alt < dist[neighbor] { // relax (u,v)
                distance[neighbor] := alt
                previous[neighbor] := smallest
            }
        }
	}
	return previous
}
```

### A* (A-star)

> A* is pretty similar to Dijkstra's algorithm, but it uses a heuristic to speed up the search.
 
###### pseudocode
```go
func AStar(Graph, source, target) {
    for _, v := range Graph.Vertices{
        distance[v] = INFINITE // initial distance from source to v is infinite
        previous[v] = UNDEFINED // previous node in optimal path from source
    }
    distance[source] := 0 // distance from source to source
    Q := // the set of all nodes in Graph
    for Q is not empty {
        smallest := //node in Q with best result from heuristic // This is the difference between A* and Dijkstra
        remove(smallest, Q)
        for _, neighbor := range smallest.neighbors { // where neighbor has not yet been removed from Q.
            alt := dist[smallest] + length(smallest, neighbor)
            if alt < dist[neighbor] { // relax (u,v)
                distance[neighbor] := alt
                previous[neighbor] := smallest
            }
        }
    }
    return previous
}
```

### Binary Search

> Binary search is a search algorithm that finds the position of a target value within a sorted array.
> It works by repeatedly dividing in half the portion of the list that could contain the item, until you've narrowed down the possible locations to just one.

###### pseudocode
```go
func binarySearch(arr, target) {
    left := 0
    right := len(arr) - 1
    
    for left <= right {
        mid := (left + right) / 2
        
		if arr[mid] == target {
            return mid
        }
		
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
	
	return -1
}
```

### Greedy

> A greedy algorithm is any algorithm that follows the problem-solving heuristic of making the locally optimal choice at each stage. 
> In many problems, a greedy strategy does not produce an optimal solution, but a greedy heuristic can yield locally optimal solutions that approximate a globally optimal solution in a reasonable amount of time.

### Backtracking

> Backtracking is an algorithmic technique where the goal is to get all solutions to a problem using the brute force approach. 
> It consists of building a set of all the solutions incrementally. Since a problem would have constraints, the solutions that fail to satisfy them will be removed.
> 
> It uses recursive calling to find a solution set by building a solution step by step, increasing levels with time. 
> In order to find these solutions, a search tree named state-space tree is used. In a state-space tree, each branch is a variable, and each level represents a solution.

## Problems

| #  | Problem | Difficulty           |
|----|---------|----------------------|
| 1  | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)| :star: :star:        |
| 2  | [ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/description/)| :star: :star:        |
| 3  | [Find All Possible Recipes From Given Supplies](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/)| :star: :star:        |
| 4  | [Remove Palindromic Subsequences](https://leetcode.com/problems/remove-palindromic-subsequences/)| :star:               |
| 5  | [Find the String with LCP](https://leetcode.com/problems/find-the-string-with-lcp/)| :star: :star: :star: |
| 6  | [Two Sum](https://leetcode.com/problems/two-sum/)| :star:               |
| 7  | [Remove Duplication From Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)| :star: :star:        |
| 8  | [Container with Most Water](https://leetcode.com/problems/container-with-most-water/)| :star: :star:        |
| 9  | [Find All Duplicates in Array](https://leetcode.com/problems/find-all-duplicates-in-an-array)| :star: :star:        |
| 10 | [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)| :star: :star:        |

---
#### :construction: Could not solve yet

- [Largest Color Value in a Directed Graph](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/)