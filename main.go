package main

import (
	"time"

	"github.com/flow-lab/dlog"
)

func main() {
	logger := dlog.NewLogger("ms")

	for true {
		logger.Info("Still alive...")
		time.Sleep(1000 * time.Millisecond)
	}
}
