package utils

type Vector struct {
	X int
	Y int
	Z int
}

func NewVector(x, y, z int) (v Vector) {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

func NewZeroVector() (v Vector) {
	v.X = 0
	v.Y = 0
	v.Z = 0
	return v
}

func VecAdd(p, q Vector) (v Vector) {
	v.X = p.X + q.X
	v.Y = p.Y + q.Y
	v.Z = p.Z + q.Z
	return v
}
