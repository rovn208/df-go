package counter

import "container/list"

type Counter interface {
	CountRectangles() int
}

type RectangleCounter struct {
	Matrix [][]int
}

func NewCounter(matrix [][]int) *RectangleCounter {
	return &RectangleCounter{Matrix: matrix}
}

func (c *RectangleCounter) CountRectangles() int {
	if len(c.Matrix) == 0 {
		return 0
	}

	rectangles := make([][]*Point, 0)
	visited := make(map[string]*Point)
	for i := 0; i < len(c.Matrix); i++ {
		for j := 0; j < len(c.Matrix[0]); j++ {
			point := NewPoint(i, j)
			if _, ok := visited[point.Id]; ok || c.Matrix[i][j] == 0 {
				continue
			}
			if rectangle := GetRectangle(c.Matrix, point, visited); len(rectangle) > 0 {
				rectangles = append(rectangles, rectangle)
			}
		}
	}

	return len(rectangles)
}

func GetRectangle(matrix [][]int, p *Point, visited map[string]*Point) []*Point {
	valid := true
	points := make([]*Point, 0)
	points = append(points, p)
	queue := list.New()
	queue.PushBack(p)
	visited[p.Id] = p

	// Get top-right node
	topRight := GetTopRight(matrix, p)

	// BFS
	for queue.Len() > 0 {
		qnode := queue.Front()
		for _, point := range GetNeighbors(matrix, qnode.Value.(*Point)) {
			if _, ok := visited[point.Id]; !ok {
				visited[point.Id] = point
				queue.PushBack(point)
				points = append(points, point)

				if !IsInRectangle(point, p, topRight) {
					valid = false
				}
			}
		}
		queue.Remove(qnode)
	}

	rectangleWidth := topRight.Col - p.Col + 1
	if valid && len(points)%rectangleWidth == 0 {
		return points
	}
	return make([]*Point, 0)

}

func GetTopRight(matrix [][]int, p *Point) *Point {
	topRight := NewPoint(p.Row, p.Col+1)
	for IsValidPoint(matrix, topRight) {
		topRight = NewPoint(p.Row, topRight.Col+1)
	}
	return NewPoint(p.Row, topRight.Col-1)
}

func GetNeighbors(matrix [][]int, point *Point) []*Point {
	rs := make([]*Point, 0)
	indices := []*Point{NewPoint(point.Row-1, point.Col), NewPoint(point.Row+1, point.Col), NewPoint(point.Row, point.Col-1), NewPoint(point.Row, point.Col+1)}
	for _, p := range indices {
		if IsValidPoint(matrix, p) {
			rs = append(rs, p)
		}
	}
	return rs
}

func IsValidPoint(arr [][]int, point *Point) bool {
	row, col := point.Row, point.Col
	return row >= 0 && row < len(arr) && col >= 0 && col < len(arr[0]) && arr[row][col] == 1
}
