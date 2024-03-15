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
	*components.Velocity
	*components.Outline
}

type RenderSystem struct {
	entities map[uint64]renderEntity
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}

func (s *RenderSystem) Priority() int {
	return 3
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
	vel *components.Velocity,
	ol *components.Outline,
) {
	s.entities[basic.ID()] = renderEntity{
		basic,
		pos,
		size,
		sc,
		sh,
		vel,
		ol,
	}
}

func (s *RenderSystem) Update(dt float32) {
	s.handleRendering(dt)
}

func (s *RenderSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *RenderSystem) handleRendering(dt float32) {
	for _, entity := range s.entities {
		switch entity.Shape.Value {
		case "Player":
			s.drawPlayer(entity, dt)
		}
	}
}

func (s *RenderSystem) drawPlayer(entity renderEntity, dt float32) {
	// Main Body
	pos := entity.Position.Vector2

	rl.DrawRectangleV(pos, entity.Size.Vector2, rl.Black)

	// Top Side Horizontal Outline
	rl.DrawRectangleV(
		pos,
		entity.Outline.HorizontalSize,
		entity.Outline.Color,
	)

	// Left Side Vertical Outline
	rl.DrawRectangleV(
		pos,
		entity.Outline.VerticalSize,
		entity.Outline.Color,
	)

	// Right Side Vertical Outline
	rl.DrawRectangleV(
		rl.Vector2Add(
			rl.Vector2Subtract(pos, entity.Outline.SX),
			rl.Vector2{X: entity.Size.X, Y: 0},
		),
		entity.Outline.VerticalSize,
		entity.Outline.Color,
	)

	// Bottom Side Horizontal Outline
	rl.DrawRectangleV(
		rl.Vector2Add(
			rl.Vector2Subtract(pos, entity.Outline.SY),
			rl.Vector2{X: 0, Y: entity.Size.Y},
		),
		entity.Outline.HorizontalSize,
		entity.Outline.Color,
	)

	// Left Eye
	eyePosScale := rl.Vector2Scale(entity.Velocity.Vector2, 5)
	rl.DrawCircleV(
		rl.Vector2Add(
			pos,
			rl.Vector2Add(rl.Vector2{X: 25, Y: 25}, eyePosScale),
		),
		5,
		entity.Outline.Color,
	)

	// Right Eye
	rl.DrawCircleV(
		rl.Vector2Add(
			rl.Vector2Add(pos, rl.Vector2{X: entity.Size.X, Y: 0}),
			rl.Vector2Add(rl.Vector2{X: -25, Y: 25}, eyePosScale),
		),
		5,
		entity.Outline.Color,
	)
}
