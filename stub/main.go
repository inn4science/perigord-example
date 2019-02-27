// Invokes the perigord driver application

package main

import (
	_ "gitlab.com/go-truffle/perigord-contract-example/migrations"
	"gitlab.com/go-truffle/enhanced-perigord/stub"
)

func main() {
	stub.StubMain()
}
