package component

import (
	"github.com/bitver/donburi"
)

type HueData struct {
	Colorful *bool
	Value    float64
}

var Hue = donburi.NewComponentType[HueData]()
