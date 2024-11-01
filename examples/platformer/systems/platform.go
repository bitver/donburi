package systems

import (
	"image/color"

	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/platformer/components"
	dresolv "github.com/bitver/donburi/examples/platformer/resolv"
	"github.com/bitver/donburi/examples/platformer/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func UpdateFloatingPlatform(ecs *ecs.ECS) {
	for e := range tags.FloatingPlatform.Iter(ecs.World) {
		tw := components.Tween.Get(e)
		// Platform movement needs to be done first to make sure there's no space between the top and the player's bottom; otherwise, an alternative might
		// be to have the platform detect to see if the Player's resting on it, and if so, move the player up manually.
		y, _, seqDone := tw.Update(1.0 / 60.0)

		obj := dresolv.GetObject(e)
		obj.Y = float64(y)
		if seqDone {
			tw.Reset()
		}
	}
}

func DrawPlatform(ecs *ecs.ECS, screen *ebiten.Image) {
	for e := range tags.Platform.Iter(ecs.World) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{180, 100, 0, 255}
		ebitenutil.DrawRect(screen, o.X, o.Y, o.W, o.H, drawColor)
	}
}

func DrawFloatingPlatform(ecs *ecs.ECS, screen *ebiten.Image) {
	for e := range tags.FloatingPlatform.Iter(ecs.World) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{180, 100, 0, 255}
		ebitenutil.DrawRect(screen, o.X, o.Y, o.W, o.H, drawColor)
	}
}
