package core

import (
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/ldtkgo"

	"github.com/ivandrenjanin/2d-game-go/constants"
)

func RunGame() {
	rl.InitWindow(
		constants.SCREEN_DEFAULT_WIDTH,
		constants.SCREEN_DEFAULT_HEIGHT,
		constants.WINDOW_TITLE,
	)
	defer rl.CloseWindow()
	// rl.SetTargetFPS(constants.SCREEN_TARGET_FPS)
	rl.SetWindowState(rl.FlagWindowResizable)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fp := path.Join(wd, "project.ldtk")

	project, err := ldtkgo.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}

	level := project.Levels[0]
	worldX := int32(level.WorldX)
	worldY := int32(level.WorldY)
	levelWidth := int32(level.Width)
	levelHeight := int32(level.Height)
	r, g, b, a := level.BGColor.RGBA()
	bgColor := rl.Color{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}

	var camRotation float32 = 0
	var camZoom float32 = 1
	cam := rl.NewCamera2D(
		rl.Vector2{
			X: float32(constants.SCREEN_DEFAULT_WIDTH) / 2,
			Y: float32(constants.SCREEN_DEFAULT_HEIGHT) / 2,
		},
		rl.Vector2{},
		camRotation,
		camZoom,
	)
	w := createWorld(&cam)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Game Logic Start
		rl.BeginMode2D(cam)

		cam.Zoom += rl.GetMouseWheelMove() * 0.05
		if cam.Zoom > 3.0 {
			cam.Zoom = 3.0
		} else if cam.Zoom < 1.0 {
			cam.Zoom = 1.0
		}
		rl.DrawRectangle(worldX, worldY, levelWidth, levelHeight, bgColor)
		w.Update(rl.GetFrameTime())
		rl.EndMode2D()
		// Game Logic End
		rl.EndDrawing()
	}
}
