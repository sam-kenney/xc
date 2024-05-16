package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"sam.kenney.blog/xc/models"
	"sam.kenney.blog/xc/exchange"
)

// Set via -ldflags="-X main.apiKey='$(EXCHANGE_API_KEY)'"
// in go toolchain.
var apiKey = ""

func main() {
	exch := models.Xc{}
	arg.MustParse(&exch)

	if err := exchange.Get(exch, apiKey); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
