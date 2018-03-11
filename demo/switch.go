package demo

import (
	"fmt"
	"runtime"
	"time"
)

func TSwitch() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X")
	case "linux":
		fmt.Print("Linux")
	default:
		fmt.Printf("%s", os)
	}

	fmt.Print("\nWHen's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Print("Today")
	case today + 1:
		fmt.Print("Tomorrow")
	case today + 2:
		fmt.Print("In two days.")
	default:
		fmt.Print("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("\n Good morning!")
	case t.Hour() < 17:
		fmt.Println("\n Good afternoon.")
	default:
		fmt.Println("\n Good evening.")
	}
}
