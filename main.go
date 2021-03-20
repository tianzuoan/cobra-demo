package main

import (
	"github.com/tianzuoan/cobra-demmo/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd execute failed! err:", err)
	}
}
