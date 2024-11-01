package filter

import "github.com/bitver/donburi/component"

type or struct {
	filters []LayoutFilter
}

func Or(filters ...LayoutFilter) LayoutFilter {
	return &or{filters: filters}
}

func (f *or) MatchesLayout(components []component.IComponentType) bool {
	for _, filter := range f.filters {
		if filter.MatchesLayout(components) {
			return true
		}
	}
	return false
}
