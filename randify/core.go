package randify

import (
	"fmt"
	"math/rand"
	"sync"
)

// Context : contains pathToFile and random gen info
type Context struct {
	pathToFile string
	numFile    int
	outFname   string
	identifier string
}

// NewContext : create a pointer to Context
func NewContext(fname string, seed int64, numFile int, outFname string, identifier string) *Context {
	c := new(Context)
	c.pathToFile = fname
	c.numFile = numFile
	c.outFname = outFname
	c.identifier = identifier
	rand.Seed(seed)
	return c
}

// Run : runs the context
func (c *Context) Run() {
	var wg sync.WaitGroup
	data := readFile(c.pathToFile)

	for i := 0; i < c.numFile; i++ {

		wg.Add(1)
		go func(fi int) {
			data = ReplaceRand(data, c.identifier)
			writeFile(c.outFname+fmt.Sprint(fi)+".randify", data)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
