package main

import (
	"testing"

	_ "trusted_market/tests"
	_ "trusted_market/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
