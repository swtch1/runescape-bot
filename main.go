package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	depositPos       = position{x: 355, y: 340}
	xPos             = position{x: 490, y: 100}
	rockInitialPos   = position{x: 310, y: 155}
	rockNewPos       = position{x: 280, y: 180}
	timesToClickRock = 3
	bankPos          = position{x: 210, y: 245}
	// number of times to run the whole process. Keep at a low number for testing.  Set to -1 to run forever.
	timesToRunTheWholeThing = 1
)

func main() {
	if timesToRunTheWholeThing < 0 {
		timesToRunTheWholeThing = 1_000_000
	}
	for i := 0; i < timesToRunTheWholeThing; i++ {
		depositPos.moveTo()
		time.Sleep(time.Millisecond * 100)
		click()
		time.Sleep(time.Second * 1)
		xPos.moveTo()
		time.Sleep(time.Millisecond * 100)
		click()
		time.Sleep(time.Second * 1)
		rockInitialPos.moveTo()
		time.Sleep(time.Millisecond * 100)
		click()
		time.Sleep(time.Second * 3)
		rockNewPos.moveTo()
		time.Sleep(time.Millisecond * 100)
		for i := 0; i < timesToClickRock; i++ {
			click()
			time.Sleep(time.Second * 3)
		}
		bankPos.moveTo()
		time.Sleep(time.Millisecond * 100)
		click()
		time.Sleep(time.Millisecond * 100)

		time.Sleep(time.Millisecond * 500)
	}
}

type position struct {
	x int
	y int
}

func (p position) moveTo() {
	robotgo.Move(p.x, p.y)
}

func click() {
	robotgo.Click()
}
