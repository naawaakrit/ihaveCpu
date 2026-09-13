// Copyright (c) 2026 Nawakarit
// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package power

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var powerDetailLabel *widget.Label //ประกาศแบบ golbal
func PowerDetailLabelcmd(text string) {
	if powerDetailLabel != nil {
		powerDetailLabel.SetText(text)
	}
}

var batteryDetailLabel *widget.Label //ประกาศแบบ golbal
func BatteryDetailLabelcmd(text string) {
	if batteryDetailLabel != nil {
		batteryDetailLabel.SetText(text)
	}
}

func PowerTabs() fyne.CanvasObject {

	powerDetailLabel = widget.NewLabel("")
	batteryDetailLabel = widget.NewLabel("")

	sub_power := container.NewVBox(
		widget.NewCard("Power", "", powerDetailLabel),
	)

	sub_battery := container.NewVBox(
		widget.NewCard("Battery", "", batteryDetailLabel),
	)

	return container.NewAppTabs(
		container.NewTabItem("Power/Detail", container.NewScroll(sub_power)),
		container.NewTabItem("Battery", container.NewScroll(sub_battery)),
	)
}
