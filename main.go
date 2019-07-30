package main

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne"

	"math/rand"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Random Button v2")
	// Champion Tab
	b := widget.NewHBox(widget.NewTabContainer(mainTab(), champTab(), posTab(), roleTab()))
	w.SetContent(b)
	w.Resize(fyne.NewSize(200, 600))

	w.ShowAndRun()
}

var selChamps = initChamps()

func initChamps() map[string]bool {
	arr := make(map[string]bool, len(champions))
	for _, v := range champions {
		arr[v.name] = true
	}
	return arr
}

func updateChamp(c string, s bool) {
	selChamps[c] = s
	fmt.Printf("Updated Champion %v; New selection is %v\n", c, selChamps)
}

var selRoles = initRoles()

func initRoles() map[string]bool {
	arr := make(map[string]bool, len(roles))
	for _, v := range roles {
		arr[v] = true
	}
	return arr
}

func updateRole(r string, s bool) {
	selRoles[r] = s
	fmt.Printf("Updated Role %v; New selection is %v\n", r, selRoles)
}

var selPos = initPos()

func initPos() map[string]bool {
	arr := make(map[string]bool, len(positions))
	for _, v := range positions {
		arr[v] = true
	}
	return arr
}

func updatePos(p string, s bool) {
	selPos[p] = s
	fmt.Printf("Updated Position %v; New selection is %v\n", p, selPos)
}

func roll(res []*widget.Label) {
	candidates := make([]champion, 0, len(champions))
	for _, v := range champions {
		if !selChamps[v.name] {
			continue
		}
		possibleP := make([]string, 0, len(v.positions))
		for _, p := range v.positions {
			if selPos[p] {
				possibleP = append(possibleP, p)
			}
		}
		if len(possibleP) == 0 {
			continue
		}
		possibleR := make([]string, 0, len(v.roles))
		for _, r := range v.roles {
			if selRoles[r] {
				possibleR = append(possibleR, r)
			}
		}
		if len(possibleR) == 0 {
			continue
		}
		candidates = append(candidates, champion{v.name, possibleP, possibleR})
	}
	if len(candidates) == 0 {
		return
	}
	c := candidates[rand.Intn(len(candidates))]
	res[0].SetText(c.name)
	p := c.positions[rand.Intn(len(c.positions))]
	r := c.roles[rand.Intn(len(c.roles))]
	res[1].SetText("Role: " + r)
	res[2].SetText("Position: " + p)
}

// Champions & Coversion

type champion struct {
	name      string
	positions []string
	roles     []string
}

func isRole(s string) int {
	for i, v := range roles {
		if strings.ToUpper(s) == strings.ToUpper(v) {
			return i
		}
	}
	return -1
}

func isPosition(s string) int {
	for i, v := range positions {
		if strings.ToUpper(s) == strings.ToUpper(v) {
			return i
		}
	}
	return -1
}

func newChamp(l string) champion {
	s := strings.Split(l, " ")
	if len(s) == 0 {
		log.Fatal("String is too short, must have a name")
	}
	n := s[0]
	p := make([]string, 0, 1)
	r := make([]string, 0, 1)
	for _, v := range s[1:] {
		x := isPosition(v)
		if x >= 0 {
			p = append(p, positions[x])
			continue
		}
		x = isRole(v)
		if x >= 0 {
			r = append(r, roles[x])
			continue
		}
	}
	return champion{n, p, r}
}
