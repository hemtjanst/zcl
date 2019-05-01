package main

import (
	"log"
	"neotor.se/zcl/generator"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	check(codegen.GenerateCluster(
		"neotor.se/zcl/cluster",
		"./definition",
		"./template",
		"./cluster",
	))

}
