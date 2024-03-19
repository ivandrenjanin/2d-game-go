package tilestacker

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/ldtkgo"
)

type Project struct {
	project *ldtkgo.Project
	Color   rl.Color
	Levels  []*Level
}

type Level struct {
	Level    *ldtkgo.Level
	Position rl.Vector2
	Size     rl.Vector2
	Layers   []*Layer
}

type Layer struct {
	Layer *ldtkgo.Layer
	Tiles []*Tile
}

type Tile struct {
	Tile     *ldtkgo.Tile
	Position rl.Vector2
	Size     rl.Vector2
}

func Open(filePath string) (*Project, error) {
	project, err := ldtkgo.Open(filePath)
	if err != nil {
		return nil, err
	}

	parsedData := transformProjectData(project)

	return parsedData, nil
}

func transformProjectData(pr *ldtkgo.Project) *Project {
	newProject := Project{}

	r, g, b, a := pr.BGColor.RGBA()
	bgColor := rl.Color{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	newProject.Color = bgColor

	lvls := make([]*Level, len(pr.Levels), cap(pr.Levels))

	for i, l := range pr.Levels {
		lvl := Level{
			Level: l,
		}
		worldX := float32(l.WorldX)
		worldY := float32(l.WorldY)
		levelWidth := float32(l.Width)
		levelHeight := float32(l.Height)

		lvl.Position = rl.Vector2{X: worldX, Y: worldY}
		lvl.Size = rl.Vector2{X: levelWidth, Y: levelHeight}

		lyrs := make([]*Layer, len(l.Layers), cap(l.Layers))

		for ii, ll := range l.Layers {
			lyr := Layer{
				Layer: ll,
			}
			ts := make([]*Tile, len(ll.Tiles), cap(ll.Tiles))
			gs := float32(ll.GridSize)
			for iii, lll := range ll.Tiles {
				t := Tile{
					Tile:     lll,
					Position: rl.Vector2{X: float32(lll.Position[0]), Y: float32(lll.Position[1])},
					Size:     rl.Vector2{X: gs, Y: gs},
				}
				ts[iii] = &t
			}

			lyr.Tiles = ts
			lyrs[ii] = &lyr
		}

		lvl.Layers = lyrs
		lvls[i] = &lvl
	}
	newProject.Levels = lvls

	return &newProject
}
