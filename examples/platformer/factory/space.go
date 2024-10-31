package factory

import (
	"github.com/bitver/donburi"
	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/platformer/archetypes"
	"github.com/bitver/donburi/examples/platformer/components"
	"github.com/bitver/donburi/examples/platformer/config"
	"github.com/solarlune/resolv"
)

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)

	cfg := config.C
	spaceData := resolv.NewSpace(cfg.Width, cfg.Height, 16, 16)
	components.Space.Set(space, spaceData)

	return space
}
