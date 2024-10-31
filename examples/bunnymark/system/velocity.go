package system

import (
	"github.com/bitver/donburi"
	"github.com/bitver/donburi/examples/bunnymark/component"
	"github.com/bitver/donburi/filter"
)

type Velocity struct {
	query *donburi.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		query: donburi.NewQuery(filter.Contains(component.Position, component.Velocity)),
	}
}

func (v *Velocity) Update(w donburi.World) {
	for entry := range v.query.Iter(w) {

		position := component.Position.Get(entry)
		velocity := component.Velocity.Get(entry)

		position.X += velocity.X
		position.Y += velocity.Y
	}
}
