package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"os"
	"strconv"
	"time"
)
func main() {
	var delayDurationInSecs = 60
	if(len(os.Args) > 1) {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			delayDurationInSecs = n
		} else {
			fmt.Println("Bad input:", os.Args[1], ". Expected a number")
			os.Exit(1)
		}
	}

	screenSizeX, screenSizeY := robotgo.GetScreenSize()
	for {
		robotgo.MoveMouseSmooth(rand.Intn(screenSizeX), rand.Intn(screenSizeY))
		time.Sleep(time.Duration(delayDurationInSecs) * time.Second)
	}
}
