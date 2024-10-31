package components

import (
	"github.com/bitver/donburi"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	SpeedX         float64
	SpeedY         float64
	OnGround       *resolv.Object
	WallSliding    *resolv.Object
	FacingRight    bool
	IgnorePlatform *resolv.Object
}

var Player = donburi.NewComponentType[PlayerData]()
