package factory

import (
	"github.com/bitver/donburi"
	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/platformer/archetypes"
	dresolv "github.com/bitver/donburi/examples/platformer/resolv"
	"github.com/solarlune/resolv"
)

func CreateWall(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	wall := archetypes.Wall.Spawn(ecs)
	dresolv.SetObject(wall, obj)
	return wall
}
