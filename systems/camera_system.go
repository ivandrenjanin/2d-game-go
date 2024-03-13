package systems

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
	"github.com/ivandrenjanin/2d-game-go/constants"
)

type cameraEntity struct {
	*ecs.BasicEntity
	*components.Position
	*components.Shape
}

type CameraSystem struct {
	entities map[uint64]cameraEntity
	Camera   rl.Camera2D
}

func NewCameraSystem() CameraSystem {
	return CameraSystem{}
}

func (s *CameraSystem) New(w *ecs.World) {
	s.entities = make(map[uint64]cameraEntity)
	s.Camera = rl.NewCamera2D(
		rl.Vector2{
			X: float32(constants.SCREEN_DEFAULT_WIDTH) / 2,
			Y: float32(constants.SCREEN_DEFAULT_HEIGHT) / 2,
		},
		rl.Vector2{},
		0,
		1,
	)
}

func (s *CameraSystem) Add(
	basic *ecs.BasicEntity,
	pos *components.Position,
	sp *components.Shape,
) {
	s.entities[basic.ID()] = cameraEntity{
		basic,
		pos,
		sp,
	}
}

func (s *CameraSystem) Update(dt float32) {
	for _, entity := range s.entities {
		switch entity.Shape.Value {
		case "Player":
			s.handleCamera(entity)
		}
	}
}

func (s *CameraSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *CameraSystem) handleCamera(entity cameraEntity) {
	s.Camera.Target = rl.Vector2Add(entity.Position.Vector2, rl.Vector2{X: 20, Y: 20})
}
