package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type debugInfoEntity struct {
	*ecs.BasicEntity
	*components.Position
	*components.Velocity
	*components.Size
	*components.Shape
}

type DebugInfoSystem struct {
	entities  map[uint64]debugInfoEntity
	debugInfo bool
}

func NewDebugInfoSystem() DebugInfoSystem {
	return DebugInfoSystem{}
}

func (s *DebugInfoSystem) New(w *ecs.World) {
	s.entities = make(map[uint64]debugInfoEntity)
}

func (s *DebugInfoSystem) Add(
	basic *ecs.BasicEntity,
	pos *components.Position,
	vel *components.Velocity,
	size *components.Size,
	sp *components.Shape,
) {
	s.entities[basic.ID()] = debugInfoEntity{
		basic,
		pos,
		vel,
		size,
		sp,
	}
}

func (s *DebugInfoSystem) Update(dt float32) {
	s.handleDebugToggle()

	for _, entity := range s.entities {
		switch entity.Shape.Value {
		case "Player":
			s.drawPlayerDebugInfo(entity)
		}
	}

	s.drawGeneralDebugInfo()
}

func (s *DebugInfoSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *DebugInfoSystem) handleDebugToggle() {
	if rl.IsKeyPressed(rl.KeyP) {
		s.debugInfo = !s.debugInfo
	}
}

func (s *DebugInfoSystem) drawGeneralDebugInfo() {
	if !s.debugInfo {
		return
	}

	fps := rl.GetFPS()

	rl.DrawRectangle(5, 5, 150, 45, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("FPS: %d", fps), 8, 8, 20, rl.Black)
}

func (s *DebugInfoSystem) drawPlayerDebugInfo(entity debugInfoEntity) {
	if !s.debugInfo {
		return
	}

	posXInt := int32(entity.Position.X)
	posYInt := int32(entity.Position.Y)
	sizeXInt := int32(entity.Size.X)
	var step int32 = 0
	var stepInc int32 = 20
	var fontSize int32 = 20
	// sizeYInt := int32(entity.Size.Y)

	// Debug Title
	rl.DrawText("Player", posXInt-5, posYInt-25, fontSize, rl.RayWhite)

	// Debug Position X, Y
	rl.DrawText(
		fmt.Sprintf("PosX: %.2f", entity.Position.X),
		posXInt+sizeXInt+5,
		posYInt+step,
		fontSize,
		rl.RayWhite,
	)
	step += stepInc

	rl.DrawText(
		fmt.Sprintf("PosY: %.2f", entity.Position.Y),
		posXInt+sizeXInt+5,
		posYInt+step,
		fontSize,
		rl.RayWhite,
	)
	step += stepInc

	// Debug Velocity X, Y
	rl.DrawText(
		fmt.Sprintf("VelX: %.2f", entity.Velocity.X),
		posXInt+sizeXInt+5,
		posYInt+step,
		fontSize,
		rl.RayWhite,
	)
	step += stepInc

	rl.DrawText(
		fmt.Sprintf("VelY: %.2f", entity.Velocity.Y),
		posXInt+sizeXInt+5,
		posYInt+step,
		fontSize,
		rl.RayWhite,
	)
	step += stepInc
}
