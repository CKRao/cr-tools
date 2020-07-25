package main

import (
	"cr-tools/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute fail, err: %+v", err)
	}
}
