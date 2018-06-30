package main

import (
	"time"
	"fmt"
	"os"
	"github.com/cloud-jeong/go-startup/cmd/kubelet/app"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewKubeletCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}