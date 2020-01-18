package main

import (
	"flag"
	"hemtjan.st/zcl/generator"
	"log"
	"os"
)

const (
	TplPath = "./template"
	DefPath = "./definition"
	GenPath = "./cluster"
	ZdoPath = "./zdo"
	TsPath  = "./cluster/all.ts"
	PkgName = "hemtjan.st/zcl"
)

func check(stage string, err error) {
	if err != nil {
		log.Fatalf("Error while generating %s: %v", stage, err)
	}
}

func main() {
	defPath := flag.String("definition-path", DefPath, "Path to yaml definitions")
	tplPath := flag.String("template-path", TplPath, "Path to templates")
	genPath := flag.String("cluster-path", GenPath, "Path for outputting generated clusters")
	tsPath := flag.String("typescript-path", TsPath, "Path for outputting generated typescript file")
	zdpPath := flag.String("zdp-path", ZdoPath, "Path for outputting generated ZDP commands")
	pkgName := flag.String("package", PkgName, "Generated package name")
	flag.Parse()

	os.MkdirAll(*genPath, 0755)
	os.MkdirAll(*zdpPath, 0755)

	check("zdo", generator.GenerateZdo(
		*pkgName,
		*defPath,
		*tplPath,
		*zdpPath,
	))

	check("cluster", generator.GenerateStruct(
		*pkgName,
		*defPath,
		*tplPath,
		*genPath,
	))

	check("typescript", generator.GenerateTypescript(
		"ZigBee",
		*defPath,
		*tplPath,
		*tsPath,
	))

}
