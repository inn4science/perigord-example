package main

import (
	"testing"

	_ "gitlab.com/go-truffle/perigord-contract-example/migrations"
	_ "gitlab.com/go-truffle/perigord-contract-example/tests"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
