package main

import (
	"log"
	"openshift-rollouter/controller"
)

func main() {
	err := controller.Init()
	if err != nil {
		log.Fatal(err)
	}
}
