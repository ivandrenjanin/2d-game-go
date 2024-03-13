package core

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/constants"
	"github.com/ivandrenjanin/2d-game-go/entities"
	"github.com/ivandrenjanin/2d-game-go/systems"
)

func RunGame() {
	rl.InitWindow(
		constants.SCREEN_DEFAULT_WIDTH,
		constants.SCREEN_DEFAULT_HEIGHT,
		constants.WINDOW_TITLE,
	)
	defer rl.CloseWindow()
	rl.SetTargetFPS(constants.SCREEN_TARGET_FPS)

	w := ecs.World{}
	pis := systems.NewPlayerInputSystem()
	pms := systems.NewPlayerMovementSystem()
	rs := systems.NewRenderSystem()
	w.AddSystem(&pis)
	w.AddSystem(&pms)
	w.AddSystem(&rs)

	p := entities.NewPlayer()

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *systems.PlayerInputSystem:
			sys.Add(&p.BasicEntity, &p.Velocity)
		case *systems.PlayerMovementSystem:
			sys.Add(&p.BasicEntity, &p.Velocity, &p.Position, &p.Speed)
		case *systems.RenderSystem:
			sys.Add(&p.BasicEntity, &p.Position, &p.Size, &p.ShapeColor)
		}
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(constants.Violet)
		rl.GetFrameTime()
		// Game Logic Start
		dt := rl.GetFrameTime()
		w.Update(dt)

		// Game Logic End
		rl.EndDrawing()
	}
}
