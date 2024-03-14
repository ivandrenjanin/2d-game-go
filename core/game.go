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

	rl.SetWindowState(rl.FlagWindowResizable)

	w, cam := createWorld()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(constants.Violet)
		rl.GetFrameTime()
		// Game Logic Start
		rl.BeginMode2D(*cam)
		rl.PushMatrix()
		rl.Translatef(0, 25*50, 0)
		rl.Rotatef(90, 1, 0, 0)
		rl.DrawGrid(100, 50)
		rl.PopMatrix()
		w.Update(rl.GetFrameTime())
		rl.EndMode2D()
		// Game Logic End
		rl.EndDrawing()
	}
}
