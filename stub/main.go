// Invokes the perigord driver application

package main

import (
	_ "trusted_market/migrations"
	"gitlab.com/go-truffle/enhanced-perigord/stub"
)

func main() {
	stub.StubMain()
}
