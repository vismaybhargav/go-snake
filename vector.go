package main

import "fmt"

type Vector2 struct {
	x, y int
}

func (v *Vector2) Add(ov *Vector2) *Vector2 {
	v.x += ov.x
	v.y += ov.y
	return v
}

func (v *Vector2) Subtract(ov *Vector2) *Vector2 {
	v.x -= ov.x
	v.y -= ov.y
	return v
}

func (v *Vector2) ToString() string {
	return "[" + fmt.Sprint(v.x) + ", " + fmt.Sprint(v.y) + "]"
}
