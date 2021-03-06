package utils

type Vector struct {
	X int
	Y int
	Z int
}

func NewVector(x, y, z int) Vector {
	return Vector{x, y, z}
}

func NewZeroVector() Vector {
	return Vector{}
}

func VecAdd(p, q Vector) Vector {
	return Vector{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

func VecMultiply(p Vector, m int) Vector {
	return Vector{m * p.X, m * p.Y, m * p.Z}
}

func VecEqual(p, q Vector) bool {
	return p.X == q.X && p.Y == q.Y && p.Z == q.Z
}
