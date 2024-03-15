package systems

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type renderEntity struct {
	*ecs.BasicEntity
	*components.Position
	*components.Size
	*components.ShapeColor
	*components.Shape
}

type RenderSystem struct {
	entities map[uint64]renderEntity
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}

func (s *RenderSystem) New(w *ecs.World) {
	s.entities = make(map[uint64]renderEntity)
}

func (s *RenderSystem) Add(
	basic *ecs.BasicEntity,
	pos *components.Position,
	size *components.Size,
	sc *components.ShapeColor,
	sh *components.Shape,
) {
	s.entities[basic.ID()] = renderEntity{
		basic,
		pos,
		size,
		sc,
		sh,
	}
}

func (s *RenderSystem) Update(dt float32) {
	s.handleRendering()
}

func (s *RenderSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *RenderSystem) handleRendering() {
	for _, entity := range s.entities {
		switch entity.Shape.Value {
		case "Player":
			s.drawPlayer(entity)
		}
	}
}

func (s *RenderSystem) drawPlayer(entity renderEntity) {
	// Main Body
	rl.DrawRectangleV(entity.Position.Vector2, entity.Size.Vector2, rl.Black)
	// Top Side Horizontal Outline
	rl.DrawRectangleV(
		entity.Position.Vector2,
		rl.Vector2{X: entity.Size.X, Y: entity.Size.Y / 10},
		rl.Lime,
	)

	// Left Side Vertical Outline
	rl.DrawRectangleV(
		entity.Position.Vector2,
		rl.Vector2{X: entity.Size.X / 10, Y: entity.Size.Y},
		rl.Lime,
	)

	// Right Side Vertical Outline
	rl.DrawRectangleV(
		rl.Vector2Add(
			rl.Vector2Subtract(entity.Position.Vector2, rl.Vector2{X: entity.Size.X / 10, Y: 0}),
			rl.Vector2{X: entity.Size.X, Y: 0},
		),
		rl.Vector2{X: entity.Size.X / 10, Y: entity.Size.Y},
		rl.Lime,
	)

	// Bottom Side Horizontal Outline
	rl.DrawRectangleV(
		rl.Vector2Add(
			rl.Vector2Subtract(entity.Position.Vector2, rl.Vector2{X: 0, Y: entity.Size.Y / 10}),
			rl.Vector2{X: 0, Y: entity.Size.Y},
		),
		rl.Vector2{X: entity.Size.X, Y: entity.Size.Y / 10},
		rl.Lime,
	)

	// Left Eye
	rl.DrawCircleV(rl.Vector2Add(entity.Position.Vector2, rl.Vector2{X: 25, Y: 25}), 5, rl.Lime)

	// Right Eye
	rl.DrawCircleV(
		rl.Vector2Add(
			rl.Vector2Add(entity.Position.Vector2, rl.Vector2{X: entity.Size.X, Y: 0}),
			rl.Vector2{X: -25, Y: 25},
		),
		5,
		rl.Lime,
	)
}
