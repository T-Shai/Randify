package randify

import (
	"fmt"
	"math/rand"
	"sync"
)

// Context : contains pathToFile and random gen info
type Context struct {
	pathToFile string
	numFile    uint
	outFname   string
	identifier string
	outExt     string
}

// NewContext : create a pointer to Context
func NewContext(fname string, seed int64, numFile uint, outFname string, exten string, identifier string) *Context {
	c := new(Context)
	c.pathToFile = fname
	c.numFile = numFile
	c.outFname = outFname
	c.identifier = identifier
	c.outExt = exten
	rand.Seed(seed)
	return c
}

// Run : runs the context
func (c *Context) Run() {
	var wg sync.WaitGroup
	data := readFile(c.pathToFile)
	var i uint
	for i = 0; i < c.numFile; i++ {

		wg.Add(1)
		go func(fi uint) {
			data = ReplaceRand(data, c.identifier)
			writeFile(c.outFname+fmt.Sprint(fi)+"."+c.outExt, data)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
