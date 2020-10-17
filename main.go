package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"./randify"
)

func main() {
	var pathToFilename string
	var seed int64
	var nFile uint
	var outputFilemane string
	var extension string
	var identifier string

	flag.StringVar(&pathToFilename, "i", "", "Path to the input filename (must be defined)")
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "seed for random values DEFAULT: current time")
	flag.UintVar(&nFile, "n", 1, "number of file to be created (all with randomly different values)")
	flag.StringVar(&outputFilemane, "o", "output", "output file's name without extension")
	flag.StringVar(&extension, "ext", "randify", "output file's extension without dot")
	flag.StringVar(&identifier, "id", "RAND", "Identifier to put in brackets before expressions")

	flag.Parse()

	if pathToFilename == "" {
		fmt.Println("ERROR : flag \"i\" not specified")
		fmt.Println("Use \"", os.Args[0], "-h\" for more information")

	}
	c := randify.NewContext(pathToFilename, seed, nFile, outputFilemane, extension, identifier)
	c.Run()
}
