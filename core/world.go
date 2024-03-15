package core

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/entities"
	"github.com/ivandrenjanin/2d-game-go/systems"
)

func createWorld(cam *rl.Camera2D) *ecs.World {
	w := ecs.World{}

	// Create Systems

	pis := systems.NewPlayerInputSystem()
	pms := systems.NewPlayerMovementSystem()
	cs := systems.NewCameraSystem(cam)
	rs := systems.NewRenderSystem()
	dis := systems.NewDebugInfoSystem(cam)

	// Add systems to world
	w.AddSystem(&pis) // 0
	w.AddSystem(&pms) // 1
	w.AddSystem(&cs)  // 2
	w.AddSystem(&rs)  // 3
	w.AddSystem(&dis) // 4

	// Create Entities
	p := entities.NewPlayer()

	// Add Entities to Systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *systems.PlayerInputSystem:
			sys.Add(p.BasicEntity, p.Velocity)
		case *systems.PlayerMovementSystem:
			sys.Add(p.BasicEntity, p.Velocity, p.Position, p.Speed)
		case *systems.RenderSystem:
			sys.Add(p.BasicEntity, p.Position, p.Size, p.ShapeColor, p.Shape, p.Velocity, p.Outline)
		case *systems.CameraSystem:
			sys.Add(p.BasicEntity, p.Position, p.Shape)
		case *systems.DebugInfoSystem:
		}
	}

	return &w
}
