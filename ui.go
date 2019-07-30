package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var resLabels = []*widget.Label{
	widget.NewLabel("YOUR PICK"),
	widget.NewLabel("YOUR ROLE"),
	widget.NewLabel("YOUR POSITION"),
}

func mainTab() *widget.TabItem {
	b := widget.NewVBox(
		widget.NewButton("LET'S ROLL", func() {
			roll(resLabels)
		}),
	)
	for _, l := range resLabels {
		b.Append(l)
	}
	return widget.NewTabItem("Main", b)
}

func champTab() *widget.TabItem {
	sels := make([]fyne.CanvasObject, len(champions)+1)
	sels[0] = widget.NewCheck("All", func(s bool) { selectAll(sels, s) })
	for i, v := range champions {
		c := v
		sels[i+1] = widget.NewCheck(v.name, func(b bool) { updateChamp(c.name, b) })
	}
	champTab := widget.NewTabItem(
		"CHAMPS",
		widget.NewGroupWithScroller(
			"Pick your champs!",
			sels...,
		),
	)
	selectAll(sels, true)
	return champTab
}

func selectAll(arr []fyne.CanvasObject, s bool) {
	for _, v := range arr {
		switch x := v.(type) {
		case *widget.Check:
			x.SetChecked(s)
		}
	}
}

func posTab() *widget.TabItem {
	sels := make([]fyne.CanvasObject, len(positions)+1)
	sels[0] = widget.NewCheck("All", func(s bool) { selectAll(sels, s) })
	for i, v := range positions {
		p := v
		sels[i+1] = widget.NewCheck(v, func(b bool) { updatePos(p, b) })
	}
	tab := widget.NewTabItem(
		"POSITIONS",
		widget.NewGroupWithScroller(
			"Pick your positions!",
			sels...,
		),
	)
	selectAll(sels, true)
	return tab
}

func roleTab() *widget.TabItem {
	sels := make([]fyne.CanvasObject, len(roles)+1)
	sels[0] = widget.NewCheck("All", func(s bool) { selectAll(sels, s) })
	for i, v := range roles {
		r := v
		sels[i+1] = widget.NewCheck(v, func(b bool) { updateRole(r, b) })
	}
	tab := widget.NewTabItem(
		"ROLES",
		widget.NewGroupWithScroller(
			"Pick your roles!",
			sels...,
		),
	)
	selectAll(sels, true)
	return tab
}
