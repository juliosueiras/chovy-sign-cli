package main

import (
	"github.com/juliosueiras/chovy-sign-cli/cmakeys"
	"github.com/juliosueiras/chovy-sign-cli/kirk"
	"os"
)

func main() {
	cmakeys.GenerateKeyStr(os.Args[1])
	kirk.KirkInit()
}
