package main

import (
	"github.com/Zygimantass/beer/cmd"
	"log"
)

func main() {
	log.Println("booting up beer-backend service")
	cmd.Execute()
}
