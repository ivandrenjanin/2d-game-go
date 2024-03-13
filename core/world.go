package core

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/entities"
	"github.com/ivandrenjanin/2d-game-go/systems"
)

func createWorld() (*ecs.World, *rl.Camera2D) {
	w := ecs.World{}

	// Create Systems
	pis := systems.NewPlayerInputSystem()
	pms := systems.NewPlayerMovementSystem()
	rs := systems.NewRenderSystem()
	dis := systems.NewDebugInfoSystem()
	cs := systems.NewCameraSystem()

	// Add systems to world
	w.AddSystem(&pis)
	w.AddSystem(&pms)
	w.AddSystem(&rs)
	w.AddSystem(&dis)
	w.AddSystem(&cs)

	// Create Entities
	p := entities.NewPlayer()

	// Add Entities to Systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *systems.PlayerInputSystem:
			sys.Add(&p.BasicEntity, &p.Velocity)
		case *systems.PlayerMovementSystem:
			sys.Add(&p.BasicEntity, &p.Velocity, &p.Position, &p.Speed)
		case *systems.RenderSystem:
			sys.Add(&p.BasicEntity, &p.Position, &p.Size, &p.ShapeColor, &p.Shape)
		case *systems.CameraSystem:
			sys.Add(&p.BasicEntity, &p.Position, &p.Shape)
		case *systems.DebugInfoSystem:
			sys.Add(&p.BasicEntity, &p.Position, &p.Velocity, &p.Size, &p.Shape)
		}
	}

	return &w, &cs.Camera
}
