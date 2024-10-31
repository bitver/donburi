package systems

import (
	"github.com/bitver/donburi/ecs"
	"github.com/bitver/donburi/examples/platformer/components"
	dresolv "github.com/bitver/donburi/examples/platformer/resolv"
)

func UpdateObjects(ecs *ecs.ECS) {
	for e := range components.Object.Iter(ecs.World) {
		obj := dresolv.GetObject(e)
		obj.Update()
	}
}
