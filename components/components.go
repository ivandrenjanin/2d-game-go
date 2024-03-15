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
	TopEdge        rl.Vector2
	LeftEdge       rl.Vector2
	BottomEdge     rl.Vector2
	RightEdge      rl.Vector2
	VerticalSize   rl.Vector2
	HorizontalSize rl.Vector2
}
