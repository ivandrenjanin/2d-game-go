package constants

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	// Color Constants
	Violet rl.Color = rl.Color{
		R: 70, G: 30, B: 82, A: 255,
	}
	SlateBlue rl.Color = rl.Color{
		R: 85, G: 108, B: 201, A: 255,
	}
)

const (
	// Screen Size Constant
	SCREEN_DEFAULT_WIDTH  int32 = 800
	SCREEN_DEFAULT_HEIGHT int32 = 600

	// Target FPS
	SCREEN_TARGET_FPS int32 = 60

	// Game Title
	WINDOW_TITLE = "wrecked-angle"
)
