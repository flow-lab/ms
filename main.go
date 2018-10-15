package main

import (
	"github.com/flow-lab/dlog"
	"time"
)

func main() {
	logger := dlog.NewLogger("ms")

	for true {
		logger.Info("Still alive...")
		time.Sleep(1000 * time.Millisecond)
	}
}
