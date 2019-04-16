package main

import (
	"testing"

	"github.com/flow-lab/dlog"
	"gotest.tools/assert"
)

func TestDoStuff(t *testing.T) {
	logger := dlog.NewLogger("ms-test")

	result := DoStuff(logger)

	assert.Assert(t, result)
}
