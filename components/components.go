package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Position struct {
	rl.Vector2
}

type Velocity struct {
	rl.Vector2
}

type Size struct {
	rl.Vector2
}

type Speed struct {
	Value float32
}

type ShapeColor struct {
	rl.Color
}
