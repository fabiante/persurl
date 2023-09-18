package main

import (
	"log"

	"github.com/fabiante/persurl/cli/cmds"
)

func main() {
	if err := cmds.Root.Execute(); err != nil {
		log.Fatal(err)
	}
}
