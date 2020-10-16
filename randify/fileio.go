package randify

import "io/ioutil"

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	exceptionChecker(err)
	return string(data)
}

func writeFile(filename string, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	exceptionChecker(err)
}
