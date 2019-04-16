package main

import (
	"time"

	"github.com/flow-lab/dlog"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := dlog.NewLogger("ms")

	for true {
		DoStuff(logger)
		time.Sleep(1000 * time.Millisecond)
	}
}

// DoStuff - I am doing stuff
func DoStuff(logger *logrus.Entry) bool {
	logger.Info("Still alive...")
	return true
}
