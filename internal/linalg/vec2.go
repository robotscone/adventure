package linalg

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func New(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Add(rhs Vec2) Vec2 {
	return Vec2{
		X: v.X + rhs.X,
		Y: v.Y + rhs.Y,
	}
}

func (v Vec2) Sub(rhs Vec2) Vec2 {
	return Vec2{
		X: v.X - rhs.X,
		Y: v.Y - rhs.Y,
	}
}

func (v Vec2) Mul(rhs float64) Vec2 {
	return Vec2{
		X: v.X * rhs,
		Y: v.Y * rhs,
	}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec2) Dot(rhs Vec2) float64 {
	return v.X*rhs.X + v.Y*rhs.Y
}

func (v Vec2) MagSq() float64 {
	// The result of the dot product of a vector with itself is the magnitude
	// squared because the dot product is the scalar projection of the lhs
	// vector onto the rhs vector which is then scaled by rhs vector's magnitude
	return v.Dot(v)
}

func (v Vec2) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

func (v Vec2) Norm() Vec2 {
	// A zero vector cannot be normalised
	if v.X == 0 && v.Y == 0 {
		return v
	}

	reciprocal := 1.0 / v.Mag()

	return v.Mul(reciprocal)
}

func (v Vec2) Reflect(normal Vec2) Vec2 {
	// We assume the normal is already a unit vector, hence the name "normal"
	// The reason for this is that if the user has already normalised the vector
	// then we would be wasting cycles re-normalising it again
	scalarProjectionDoubled := v.Dot(normal) * 2.0
	normalScaled := normal.Mul(scalarProjectionDoubled)

	return normalScaled.Sub(v)
}
