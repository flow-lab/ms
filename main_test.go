package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestDoStuff(t *testing.T) {
	result := DoStuff()

	assert.Assert(t, result)
}
