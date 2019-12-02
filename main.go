package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		DoStuff()
		time.Sleep(1000 * time.Millisecond)
	}
}

// DoStuff - I am doing stuff
func DoStuff() bool {
	fmt.Println("Still alive...")
	return true
}
