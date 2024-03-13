package systems

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type playerInputEntity struct {
	*ecs.BasicEntity
	*components.Velocity
}

type PlayerInputSystem struct {
	entities map[uint64]playerInputEntity
}

func NewPlayerInputSystem() PlayerInputSystem {
	return PlayerInputSystem{}
}

func (s *PlayerInputSystem) New(w *ecs.World) {
	s.entities = make(map[uint64]playerInputEntity)
}

func (s *PlayerInputSystem) Add(basic *ecs.BasicEntity, vel *components.Velocity) {
	s.entities[basic.ID()] = playerInputEntity{
		basic,
		vel,
	}
}

func (s *PlayerInputSystem) Update(dt float32) {
	s.handleMovement()
}

func (s *PlayerInputSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *PlayerInputSystem) handleMovement() {
	for _, e := range s.entities {
		v := e.Velocity.Vector2

		if rl.IsKeyDown(rl.KeyW) {
			v.Y -= 1
		} else if rl.IsKeyDown(rl.KeyS) {
			v.Y += 1
		} else {
			v.Y = 0
		}

		if rl.IsKeyDown(rl.KeyA) {
			v.X -= 1
		} else if rl.IsKeyDown(rl.KeyD) {
			v.X += 1
		} else {
			v.X = 0
		}

		if rl.Vector2Length(v) > 0 {
			v = rl.Vector2Normalize(v)
		}

		e.Velocity.Vector2 = v
	}
}
