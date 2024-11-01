package system

import (
	"github.com/bitver/donburi"
	"github.com/bitver/donburi/examples/bunnymark/component"
	"github.com/bitver/donburi/filter"
)

type Gravity struct {
	query *donburi.Query
}

func NewGravity() *Gravity {
	return &Gravity{
		query: donburi.NewQuery(filter.Contains(component.Velocity, component.Gravity)),
	}
}

func (g *Gravity) Update(w donburi.World) {
	for entry := range g.query.Iter(w) {
		gravity := component.Gravity.Get(entry)
		velocity := component.Velocity.Get(entry)

		velocity.Y += gravity.Value
	}
}
