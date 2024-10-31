package factory

import (
	"github.com/bitver/donburi"
	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/platformer/archetypes"
	"github.com/bitver/donburi/examples/platformer/components"
	dresolv "github.com/bitver/donburi/examples/platformer/resolv"
	"github.com/solarlune/resolv"
)

func CreatePlayer(ecs *ecs.ECS) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

	obj := resolv.NewObject(32, 128, 16, 24)
	dresolv.SetObject(player, obj)
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 16, 24))

	return player
}
