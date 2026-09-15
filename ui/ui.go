// Copyright (c) 2026 Nawakarit
// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package ui

import (
	"embed"
	"fmt"
	biosinfo "ihavecpu/bios"
	cpuinfo "ihavecpu/cpu"
	"ihavecpu/hardware"
	mainboardinfo "ihavecpu/mainboard"
	pcie "ihavecpu/pcie"
	pcieinfo "ihavecpu/pcie"
	powerinfo "ihavecpu/power"
	raminfo "ihavecpu/ram"
	systeminfo "ihavecpu/system"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// โหลด icon
func loadIcon(size int) fyne.Resource {
	var file string

	switch {
	case size >= 512:
		file = "assets/icons/icon-512.png" ///ที่อยู่
	case size >= 256:
		file = "assets/icons/icon-256.png"
	case size >= 128:
		file = "assets/icons/icon-128.png"
	default:
		file = "assets/icons/icon-64.png"
	}

	data, _ := iconFS.ReadFile(file)
	return fyne.NewStaticResource(file, data)
}

//go:embed assets/icons/*
var iconFS embed.FS

//go:embed assets/font/Itim-Regular.ttf
var fontItim []byte
var myFont = fyne.NewStaticResource("Itim-Regular.ttf", fontItim)

func CreateWindow() {

	a := app.NewWithID("com.nawakarit.iHaveCPU")
	a.Settings().SetTheme(&MyTheme{})
	icon := loadIcon(64)
	w := a.NewWindow("iHaveCPU")
	w.SetIcon(icon)

	pcie.InitPCI()

	system := systeminfo.SystemTabs()
	biOsTabs := biosinfo.BiosTabs()
	cpuTabs := cpuinfo.CpuTabs(w)
	mainboardTabs := mainboardinfo.MainboardTabs()
	ramTabs := raminfo.RamTabs()
	pcieTabs := pcieinfo.PcieTabs()
	powerTabs := powerinfo.PowerTabs()

	/*
		teXt, err := raminfo.GetMemoryInfo()

		if err != nil {
			teXt = err.Error()
		}
		fyne.Do(func() {
			raminfo.RamDetailLabelcmd(teXt)
		})
	*/

	//MemoryPkexec.SetText(teXt) //ให้มันอัพเดท
	/*
		cmd := exec.Command("pkexec", "bash", "-t", script)
		output, err := cmd.CombinedOutput()
		text := string(output)
		if err != nil {
			text = fmt.Sprintf("%s\n%s", text, err.Error())
			fmt.Printf("failed to run pkexec: %v\n%s\n", err, string(output))
		}
		mainboardinfo.SetMainboardPkexecAllText(text)
	*/
	tabs := container.NewAppTabs(
		container.NewTabItem("System", system),
		container.NewTabItem("Bios", biOsTabs),
		container.NewTabItem("CPU", container.NewScroll(cpuTabs)),
		container.NewTabItem("Ram", ramTabs),
		container.NewTabItem("MainBoard", container.NewScroll(mainboardTabs)),
		container.NewTabItem("Pcie", container.NewScroll(pcieTabs)),
		container.NewTabItem("Power", container.NewScroll(powerTabs)),
		//container.NewTabItem("Virtualization", container.NewScroll(nil)),
	)

	lastUpdated := widget.NewLabel("อัปเดตล่าสุด: ยังไม่มีข้อมูล")
	refreshButton := widget.NewButton("Refresh", nil)
	var refreshMu sync.Mutex
	loading := false

	applyData := func(data hardware.Data) {
		systeminfo.SystemsDetailLabelcmd(data.System)
		biosinfo.BiosDetailLabelcmd(data.BIOS)
		cpuinfo.CPUDetailLabelcmd(data.CPU)
		cpuinfo.CacheLabelcmd(data.Cache)
		raminfo.RamDetailLabelcmd(data.RAM)
		mainboardinfo.MainboardDetailLabelcmd(data.Mainboard)
		pcieinfo.PcieDetailLabelcmd(data.PCIe)
		powerinfo.PowerDetailLabelcmd(data.Power)
		powerinfo.BatteryDetailLabelcmd(data.Battery)
	}

	refreshData := func() {
		refreshMu.Lock()
		if loading {
			refreshMu.Unlock()
			return
		}
		loading = true
		refreshMu.Unlock()

		refreshButton.Disable()
		lastUpdated.SetText("กำลังอัปเดตข้อมูล...")
		go func() {
			data, err := hardware.Load()
			fyne.Do(func() {
				defer func() {
					refreshMu.Lock()
					loading = false
					refreshMu.Unlock()
					refreshButton.Enable()
				}()

				if err != nil {
					lastUpdated.SetText("อัปเดตไม่สำเร็จ")
					dialog.ShowError(err, w)
					return
				}
				applyData(data)
				lastUpdated.SetText(fmt.Sprintf("อัปเดตล่าสุด: %s", time.Now().Format("2006-01-02 15:04:05")))
			})
		}()
	}
	refreshButton.OnTapped = refreshData

	toolbar := container.NewBorder(nil, nil, refreshButton, nil, lastUpdated)
	w.SetContent(container.NewBorder(toolbar, nil, nil, nil, tabs))
	w.Resize(fyne.NewSize(720, 800))
	w.Show()
	refreshData()
	w.ShowAndRun()
}
