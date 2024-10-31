package filter

import "github.com/bitver/donburi/component"

func containsComponent(components []component.IComponentType, c component.IComponentType) bool {
	for _, component := range components {
		if component == c {
			return true
		}
	}
	return false
}
