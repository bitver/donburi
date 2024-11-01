package system

import (
	"image"

	"github.com/bitver/donburi"
	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/bunnymark_ecs/component"
	"github.com/bitver/donburi/examples/bunnymark_ecs/helper"
	"github.com/bitver/donburi/filter"
)

type bounce struct {
	bounds *image.Rectangle
	query  *donburi.Query
}

func NewBounce(bounds *image.Rectangle) *bounce {
	return &bounce{
		bounds: bounds,
		query: donburi.NewQuery(
			filter.Contains(
				component.Position,
				component.Velocity,
				component.Sprite,
			)),
	}
}

func (b *bounce) Update(ecs *ecs.ECS) {
	for entry := range b.query.Iter(ecs.World) {
		position := component.Position.Get(entry)
		velocity := component.Velocity.Get(entry)
		sprite := component.Sprite.Get(entry)

		sw, sh := float64(b.bounds.Dx()), float64(b.bounds.Dy())
		iw, ih := float64(sprite.Image.Bounds().Dx()), float64(sprite.Image.Bounds().Dy())
		relW, relH := iw/sw, ih/sh
		if position.X+relW > 1 {
			velocity.X *= -1
			position.X = 1 - relW
		}
		if position.X < 0 {
			velocity.X *= -1
			position.X = 0
		}
		if position.Y+relH > 1 {
			velocity.Y *= -0.85
			position.Y = 1 - relH
			if helper.Chance(0.5) {
				velocity.Y -= helper.RangeFloat(0, 0.009)
			}
		}
		if position.Y < 0 {
			velocity.Y = 0
			position.Y = 0
		}
	}
}
