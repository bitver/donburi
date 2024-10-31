package components

import "github.com/bitver/donburi"

type SettingsData struct {
	Debug        bool
	ShowHelpText bool
}

var Settings = donburi.NewComponentType[SettingsData]()
