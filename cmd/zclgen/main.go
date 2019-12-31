package main

import (
	"hemtjan.st/zcl/generator"
	"log"
	"os"
)

const (
	GenPath = "./cluster"
	ZdoPath = GenPath + "/zdo"
)

func check(stage string, err error) {
	if err != nil {
		log.Fatalf("Error while generating %s: %v", stage, err)
	}
}

func main() {

	os.MkdirAll(GenPath, 0755)
	os.MkdirAll(ZdoPath, 0755)

	check("zdo", generator.GenerateZdo(
		"hemtjan.st/zcl",
		"./definition",
		"./template",
		ZdoPath,
	))

	check("cluster", generator.GenerateStruct(
		"hemtjan.st/zcl",
		"./definition",
		"./template",
		GenPath,
	))

}
