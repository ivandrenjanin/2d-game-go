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
}

func NewPlayer() Player {
	pos := components.Position{Vector2: rl.Vector2{X: 0, Y: 0}}
	size := components.Size{Vector2: rl.Vector2{X: 64, Y: 64}}
	sc := components.ShapeColor{Color: rl.Red}
	spd := components.Speed{Value: 400}
	sp := components.Shape{Value: "Player"}
	be := ecs.NewBasic()
	vel := components.Velocity{}

	return Player{
		BasicEntity: &be,
		ShapeColor:  &sc,
		Speed:       &spd,
		Shape:       &sp,
		Position:    &pos,
		Size:        &size,
		Velocity:    &vel,
	}
}
