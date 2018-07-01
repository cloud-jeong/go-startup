package main

import (
	"time"
	"fmt"
	"os"
	"math/rand"
	"github.com/spf13/pflag"
	"flag"
	"github.com/cloud-jeong/go-startup/cmd/kcem/app/cmd"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Hidden flags when --help
	pflag.CommandLine.MarkHidden("version")

	command := cmd.NewKcemCommand(os.Stdin, os.Stdout, os.Stderr)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}