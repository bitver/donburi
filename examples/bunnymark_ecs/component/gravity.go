package component

import (
	"github.com/bitver/donburi"
)

type GravityData struct {
	Value float64
}

var Gravity = donburi.NewComponentType[GravityData]()
