package component

import (
	"github.com/bitver/donburi"
)

type PositionData struct {
	X, Y float64
}

var Position = donburi.NewComponentType[PositionData]()
