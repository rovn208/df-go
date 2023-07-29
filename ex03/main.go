package main

import (
	"container/list"
	"fmt"
)

// countRectangles returns the number of rectangles filled with 1s.

func countRectangles(arr [][]int) int {
	count := 0
	rectangles := make([][]*Point, 0)
	visited := make(map[string]*Point)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			point := NewPoint(i, j)
			if _, ok := visited[point.Id]; ok || arr[i][j] == 0 {
				continue
			}
			if rectangle := bfs(arr, point, visited); len(rectangle) > 0 {
				rectangles = append(rectangles, rectangle)
			}
		}
	}

	fmt.Printf("rectangles: %v\n", len(rectangles))
	for _, rectangle := range rectangles {
		if valid := checkRectangle(rectangle); valid {
			count++
		}
	}
	return count
}

func bfs(arr [][]int, p *Point, visited map[string]*Point) []*Point {
	points := make([]*Point, 0)
	queue := list.New()
	queue.PushBack(p)
	points = append(points, p)
	visited[p.Id] = p

	for queue.Len() > 0 {
		qnode := queue.Front()
		for _, point := range neighbors(arr, qnode.Value.(*Point)) {
			if _, ok := visited[point.Id]; !ok {
				visited[point.Id] = point
				queue.PushBack(point)
				points = append(points, point)
			}
		}
		queue.Remove(qnode)
	}

	return points
}

func neighbors(arr [][]int, point *Point) []*Point {
	rs := make([]*Point, 0)
	indices := []*Point{NewPoint(point.Row-1, point.Col), NewPoint(point.Row+1, point.Col), NewPoint(point.Row, point.Col-1), NewPoint(point.Row, point.Col+1)}
	for _, p := range indices {
		if isValid(arr, p) {
			rs = append(rs, p)
		}
	}
	return rs
}

func isValid(arr [][]int, point *Point) bool {
	row, col := point.Row, point.Col
	return row >= 0 && row < len(arr) && col >= 0 && col < len(arr[0]) && arr[row][col] == 1
}

func checkRectangle(points []*Point) bool {
	if len(points) == 1 {
		return true
	}
	pointMap := make(map[int]int)
	for _, p := range points {
		pointMap[p.Row]++
		pointMap[p.Col]++
	}

	for _, point := range pointMap {
		if point == len(points) {
			return true
		}
	}

	fmt.Println(points, len(points), pointMap)

	return false
}

func main() {
	arr := [][]int{
		{1, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 1},
	}

	count := countRectangles(arr)
	fmt.Printf("count: %v", count) //6
}
