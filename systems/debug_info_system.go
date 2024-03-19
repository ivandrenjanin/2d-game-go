package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type debugInfoEntity struct{}

type DebugInfoSystem struct {
	debugInfo bool
	Camera    *rl.Camera2D
}

func NewDebugInfoSystem(camera *rl.Camera2D) DebugInfoSystem {
	return DebugInfoSystem{
		Camera: camera,
	}
}

func (s *DebugInfoSystem) Priority() int {
	return 3
}

func (s *DebugInfoSystem) Add() {
}

func (s *DebugInfoSystem) Update(dt float32) {
	s.handleDebugToggle()
	s.drawDebugInfoConsole(dt)
}

func (s *DebugInfoSystem) Remove(basic ecs.BasicEntity) {
}

func (s *DebugInfoSystem) handleDebugToggle() {
	if rl.IsKeyPressed(rl.KeyP) {
		s.debugInfo = !s.debugInfo
	}
}

func (s *DebugInfoSystem) drawDebugInfoConsole(dt float32) {
	if !s.debugInfo {
		return
	}

	fps := rl.GetFPS()
	pos := rl.Vector2Subtract(
		rl.Vector2{X: s.Camera.Target.X, Y: s.Camera.Target.Y},
		rl.Vector2{X: s.Camera.Offset.X, Y: s.Camera.Offset.Y},
	)

	rl.DrawRectangleV(pos, rl.Vector2{X: 150, Y: 50}, rl.RayWhite)
	rl.DrawTextEx(
		rl.GetFontDefault(),
		fmt.Sprintf("FPS: %d", fps),
		rl.Vector2Add(pos, rl.Vector2{X: 5, Y: 5}),
		25,
		1,
		rl.Black,
	)
}
