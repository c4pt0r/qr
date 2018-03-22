package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mdp/qrterminal"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	qrterminal.Generate(string(buf), qrterminal.H, os.Stdout)
}
