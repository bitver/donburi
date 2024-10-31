package component

import (
	"github.com/bitver/donburi"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteData struct {
	Image *ebiten.Image
}

var Sprite = donburi.NewComponentType[SpriteData]()
