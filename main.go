package main

import (
	"gitlab/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
