package main

import (
	"fbi/quorum/client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
