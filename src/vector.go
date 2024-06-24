package main

import "fmt"

type Vector2 struct {
	x, y int
}

func (vector *Vector2) Add(otherVector *Vector2) *Vector2 {
	vector.x += otherVector.x
	vector.y += otherVector.y
	return vector
}

func (vector *Vector2) Subtract(otherVector *Vector2) *Vector2 {
	vector.x -= otherVector.x
	vector.y -= otherVector.y
	return vector
}

func (vector *Vector2) ToString() string {
	return "[" + fmt.Sprint(vector.x) + ", " + fmt.Sprint(vector.y) + "]"
}
