package component

import (
	"time"

	"github.com/bitver/donburi"
	"github.com/bitver/donburi/examples/bunnymark/helper"
	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsData struct {
	Ticker   *time.Ticker
	Sprite   *ebiten.Image
	Colorful bool
	Amount   int
	Gpu      string
	Tps      *helper.Plot
	Fps      *helper.Plot
	Objects  *helper.Plot
}

var Settings = donburi.NewComponentType[SettingsData]()
