package systems

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/ivandrenjanin/2d-game-go/components"
)

type playerMovementEntity struct {
	*ecs.BasicEntity
	*components.Velocity
	*components.Position
	*components.Speed
}

type PlayerMovementSystem struct {
	entities map[uint64]playerMovementEntity
}

func NewPlayerMovementSystem() PlayerMovementSystem {
	return PlayerMovementSystem{}
}

func (s *PlayerMovementSystem) New(w *ecs.World) {
	s.entities = make(map[uint64]playerMovementEntity)
}

func (s *PlayerMovementSystem) Add(
	basic *ecs.BasicEntity,
	vel *components.Velocity,
	pos *components.Position,
	spd *components.Speed,
) {
	s.entities[basic.ID()] = playerMovementEntity{
		basic,
		vel,
		pos,
		spd,
	}
}

func (s *PlayerMovementSystem) Update(dt float32) {
	s.handleMovement(dt)
}

func (s *PlayerMovementSystem) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}

func (s *PlayerMovementSystem) handleMovement(dt float32) {
	for _, entity := range s.entities {
		pos := entity.Position.Vector2
		vel := entity.Velocity.Vector2
		spd := entity.Speed.Value

		pos = rl.Vector2Add(pos, rl.Vector2Scale(vel, spd*dt))

		entity.Position.Vector2 = pos
	}
}
