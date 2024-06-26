package main

import (
	"github.com/bl-solutions/gitlab-cli/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
