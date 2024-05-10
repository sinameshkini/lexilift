package main

import (
	"lexilift/cmd"
	"log"
)

//
//var (
//	data = []string{
//		"establish",
//		"determine",
//		"identical",
//		"consumption",
//		"represent",
//		"tremendous",
//		"reveal",
//	}
//)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
