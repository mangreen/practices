package pkg

import "strconv"

type Point struct {
	// coordinate
	X int
	Y int

	// Heuristic Estimate formula：
	// F(n) = G(n) + H(n)
	// n：Current Point
	// G(n)：value from START point to current point
	// H(n)：value/distance from current to GOAL
	F int
	G int
	H int

	// IsObstacle bool

	// prev point/step
	Parent *Point
}

func NewPoint(x, y, G int, goal, parent *Point) *Point {
	p := &Point{
		X:      x,
		Y:      y,
		G:      G,
		Parent: parent,
	}

	p.H = p.CalH(goal)
	p.F = p.CalF()

	return p
}

func (p *Point) GetNeighbors(maze [][]int) map[string][]int {
	neighbors := make(map[string][]int)

	f := func(x, y int) {
		// if x y are in the maze && it is not Obstacle, then add into array
		if 0 <= x && x <= len(maze)-1 && 0 <= y && y <= len(maze[0])-1 && maze[x][y] != 1 {
			neighbors[strconv.Itoa(x)+"_"+strconv.Itoa(y)] = []int{x, y}
		}
	}

	f(p.X-1, p.Y)
	f(p.X, p.Y-1)
	f(p.X+1, p.Y)
	f(p.X, p.Y+1)

	return neighbors
}

func (p *Point) CalH(goal *Point) int {
	return abs(p.X, goal.X) + abs(p.Y, goal.Y)
}

func (p *Point) CalF() int {
	return p.G + p.H
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - 1
}
