package entities

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type Player struct {
	ecs.BasicEntity
	components.Position
	components.Velocity
	components.Speed
	components.Size
	components.ShapeColor
	components.Shape
}

func NewPlayer() Player {
	return Player{
		BasicEntity: ecs.NewBasic(),
		ShapeColor:  components.ShapeColor{Color: rl.Red},
		Speed:       components.Speed{Value: 400},
		Size:        components.Size{Vector2: rl.Vector2{X: 64, Y: 64}},
		Shape:       components.Shape{Value: "Player"},
	}
}
