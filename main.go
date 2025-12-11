package main

import (
	"lexilift/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
