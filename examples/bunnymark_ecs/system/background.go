package system

import (
	"image/color"

	"github.com/bitver/donburi/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 41, G: 44, B: 45, A: 255})
}
