package main

import (
	"fmt"

	"./randify"
)

func main() {
	text := "I am [RAND int 20 30 ] yo And i want [RAND int31 5 ] sweets ! i have exactly [RAND float32 100]€ "
	c := randify.NewContext("test", 15468)
	fmt.Println(randify.ReplaceRand(text, "RAND"))

	fmt.Println(c)
}
