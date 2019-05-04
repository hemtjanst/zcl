package main

import (
	"hemtjan.st/zcl/generator"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	check(generator.GenerateCluster(
		"hemtjan.st/zcl",
		"./definition",
		"./template",
		"./cluster",
	))

}
