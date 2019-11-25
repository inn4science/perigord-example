package main

import (
	"testing"

	_ "github.com/inn4science/perigord-example/migrations"
	_ "github.com/inn4science/perigord-example/tests"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
