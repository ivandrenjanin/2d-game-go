package entities

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type Player struct {
	*ecs.BasicEntity
	*components.Position
	*components.Velocity
	*components.Speed
	*components.Size
	*components.ShapeColor
	*components.Shape
	*components.Outline
}

func NewPlayer() Player {
	pos := components.Position{Vector2: rl.Vector2{X: 0, Y: 0}}
	size := components.Size{Vector2: rl.Vector2{X: 64, Y: 64}}
	sc := components.ShapeColor{Color: rl.Red}
	spd := components.Speed{Value: 400}
	sp := components.Shape{Value: "Player"}
	be := ecs.NewBasic()
	vel := components.Velocity{}

	sx := size.X / 10
	sy := size.Y / 10
	verticalSize := rl.Vector2{X: sx, Y: size.Y}
	horizontalSize := rl.Vector2{X: size.X, Y: sy}
	cl := rl.Lime

	ol := components.Outline{
		HorizontalSize: horizontalSize,
		VerticalSize:   verticalSize,
		ShapeColor:     components.ShapeColor{Color: cl},
		SX:             rl.Vector2{X: sx, Y: 0},
		SY:             rl.Vector2{X: 0, Y: sy},
	}

	return Player{
		BasicEntity: &be,
		ShapeColor:  &sc,
		Speed:       &spd,
		Shape:       &sp,
		Position:    &pos,
		Size:        &size,
		Velocity:    &vel,
		Outline:     &ol,
	}
}
