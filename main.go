package main

import (
	"fmt"
	"time"

	"./randify"
)

func main() {
	//text := "I am [RAND int 20 30 ] yo And i want [RAND int31 55 ] sweets ! i have exactly [RAND float32 100]€ "
	c := randify.NewContext("input.txt", time.Now().UnixNano(), 4, "output", "RAND")
	c.Run()
	fmt.Println(c)
}
