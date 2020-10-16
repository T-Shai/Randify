package randify

import "fmt"

func exceptionChecker(e error) {
	if e != nil {
		panic(e)
	}
}

func positiveChecker(expres string, i int) {
	if i < 0 {
		panic("Negative Value : " + fmt.Sprint(i) + " can't be negative in expression " + expres + " !")
	}
}
