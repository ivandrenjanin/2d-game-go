package core

import (
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/constants"
	tilestacker "github.com/ivandrenjanin/2d-game-go/tile_stacker"
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

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fp := path.Join(wd, "project.ldtk")

	project, err := tilestacker.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}

	cam := rl.NewCamera2D(
		rl.Vector2{
			X: float32(constants.SCREEN_DEFAULT_WIDTH) / 2,
			Y: float32(constants.SCREEN_DEFAULT_HEIGHT) / 2,
		},
		rl.Vector2{},
		0,
		1,
	)

	w := createWorld(&cam, project)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Game Logic Start
		rl.BeginMode2D(cam)
		w.Update(rl.GetFrameTime())
		rl.EndMode2D()
		// Game Logic End
		rl.EndDrawing()
	}
}
