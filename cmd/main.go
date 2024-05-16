package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"sam.kenney.blog/xc/exchange"
	"sam.kenney.blog/xc/models"
)

// Set via -ldflags="-X main.apiKey='$(EXCHANGE_API_KEY)'"
// in go toolchain.
var apiKey = ""

func main() {
	xc := models.Xc{}
	arg.MustParse(&xc)

	if err := exchange.Get(xc, apiKey); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
