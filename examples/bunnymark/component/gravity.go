package component

import (
	"github.com/bitver/donburi"
)

type GravityData struct {
	Value float64
}

var Gravity = donburi.NewComponentType[GravityData]()

func GetGravity(entry *donburi.Entry) *GravityData {
	return donburi.Get[GravityData](entry, Gravity)
}
