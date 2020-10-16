package randify

import "io/ioutil"

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	exceptionChecker(err)
	return string(data)
}
