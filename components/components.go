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

type Shape struct {
	Value string
}

type Outline struct {
	ShapeColor
	VerticalSize   rl.Vector2
	HorizontalSize rl.Vector2
	SX             rl.Vector2
	SY             rl.Vector2
}
