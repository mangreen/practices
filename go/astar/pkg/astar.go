package pkg

import (
	"strconv"
)

func AStar(start, goal Point, maze [][]int) *Point {
	open := new(OpenList)
	open.Init()

	start.H = start.CalH(&goal)
	start.F = start.CalF()
	open.Push(&start)

	close := map[string]*Point{}

	for len(*open.Slice) > 0 {
		// pop p has min F from open list
		p := open.Pop()

		// add p into close list
		close[strconv.Itoa(p.X)+"_"+strconv.Itoa(p.Y)] = p
		maze[p.X][p.Y] = 2

		// go through p's neighbors
		for k, v := range p.GetNeighbors(maze) {
			// skip, if v is in close list
			if _, ok := close[k]; ok {
				continue
			}

			// if v is in open list
			if n := open.IndexOf(v[0], v[1]); n != nil {
				// calculate new G
				newG := p.G + 10

				if newG < n.G {
					n.G = newG
					n.F = p.CalF()
					n.Parent = p
				}

				continue
			}

			// if v is not in open list
			n := NewPoint(v[0], v[1], p.G+10, &goal, p)

			open.Push(n)
		}

		// if goal is in open list, end
		if n := open.IndexOf(goal.X, goal.Y); n != nil {
			maze[n.X][n.Y] = 2
			return n
		}
	}

	return nil
}
