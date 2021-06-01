package main

import (
	"fbi/eth/client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
