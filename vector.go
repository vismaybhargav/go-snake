package main

type Vector2 struct {
	x, y int
}

func (v *Vector2) Add(ov *Vector2) *Vector2 {
	v.x += ov.x
	v.y += ov.y
	return v
}
