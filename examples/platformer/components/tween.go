package components

import (
	"github.com/bitver/donburi"
	"github.com/tanema/gween"
)

var Tween = donburi.NewComponentType[gween.Sequence]()
