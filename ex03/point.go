package main

import "fmt"

type Point struct {
	Id  string
	Row int
	Col int
}

func (p *Point) String() string {
	return fmt.Sprintf("%s:%d:%d", p.Id, p.Row, p.Col)
}

func NewPoint(row, col int) *Point {
	return &Point{
		Row: row,
		Col: col,
		Id:  fmt.Sprintf("%d%d", row, col),
	}
}
