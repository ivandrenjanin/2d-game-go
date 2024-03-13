package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/constants"
)

func RunGame() {
	rl.InitWindow(
		constants.SCREEN_DEFAULT_WIDTH,
		constants.SCREEN_DEFAULT_HEIGHT,
		constants.WINDOW_TITLE,
	)
	defer rl.CloseWindow()
	rl.SetTargetFPS(constants.SCREEN_TARGET_FPS)

	w := createWorld()

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
