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
		case "Rect":
			rl.DrawRectangleV(entity.Position.Vector2, entity.Size.Vector2, entity.ShapeColor.Color)
		}
	}
}
