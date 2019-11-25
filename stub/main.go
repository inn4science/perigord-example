// Invokes the perigord driver application

package main

import (
	_ "github.com/inn4science/perigord-example/migrations"
	"github.com/inn4science/perigord/stub"
)

func main() {
	stub.StubMain()
}
