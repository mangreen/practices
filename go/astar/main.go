package main

import (
	"astar/pkg"
	"fmt"
)

func main() {
	maze := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	start := pkg.Point{X: 0, Y: 0}
	goal := pkg.Point{X: 4, Y: 4}

	end := pkg.AStar(start, goal, maze)

	for end != nil {
		fmt.Print("[", end.X, ", ", end.Y, "]")

		end = end.Parent
	}

	fmt.Println()
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(maze[i][j])

			if j == len(maze[i])-1 {
				fmt.Println()
			}
		}
	}
}
