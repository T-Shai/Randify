package randify

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// ReplaceRand : parse and replace errors : interact with user when parse error otherwide panic
func ReplaceRand(data string, indenti string) string {
	//prevData := ""
	replacedData := data
	re := regexp.MustCompile(`\[([^]]+)\]`)
	matches := re.FindAllString(replacedData, -1)
	//fmt.Println(matches)
	for _, expr := range matches {
		// fmt.Println(expr)
		ri := expr
		match := strings.Trim(expr, "[")
		match = strings.Trim(match, "]")
		match = strings.TrimLeft(match, " ")
		match = strings.TrimRight(match, " ")
		if strings.HasPrefix(match, indenti) {
			args := strings.Split(match, " ")[1:]
			switch len(args) {
			case 1: // [<indentifier> <arg>]
				// fmt.Println(args[0])
				switch strings.ToLower(args[0]) {
				case "int":
					ri = fmt.Sprint(rand.Int())
					break

				case "int31":
					ri = fmt.Sprint(rand.Int31())
					break

				case "int63":
					ri = fmt.Sprint(rand.Int63())
					break

				case "uint":
					ri = fmt.Sprint(rand.Uint32())
					break

				case "uint32":
					ri = fmt.Sprint(rand.Uint32())
					break

				case "uint64":
					ri = fmt.Sprint(rand.Uint64())
					break

				case "float32":
					ri = fmt.Sprint(rand.Float32())
					break

				case "float":
					ri = fmt.Sprint(rand.Float32())
					break

				case "float64":
					ri = fmt.Sprint(rand.Float64())
					break

				case "expFloat64":
					ri = fmt.Sprint(rand.ExpFloat64())
					break

				case "normFloat64":
					ri = fmt.Sprint(rand.NormFloat64())
					break

				default:
					fmt.Println("Couldn't resolve expression :", expr)
					fmt.Println("Do you wish to ignore ? y/n")
					var ans string
					_, e := fmt.Scan(&ans)
					exceptionChecker(e)
					ans = strings.ToLower(ans)
					if ans == "y" {
						continue
					} else {
						panic("Randify : Stopped by user")
					}
				} // switch args[0]
				break // case 1

			case 2: // [<indentifier> <arg1> <arg2>]
				switch strings.ToLower(args[0]) {
				case "int":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Intn(param))
					break

				case "int31":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Int31n(int32(param)))
					break

				case "int63":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Int63n(int64(param)))
					break

				case "uint":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Uint32() % uint32(param))
					break

				case "uint32":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Uint32() % uint32(param))
					break

				case "uint64":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Uint64() % uint64(param))
					break

				case "float":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Float32() * float32(param))
					break

				case "float32":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Float32() * float32(param))
					break

				case "float64":
					param, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					positiveChecker(expr, param)
					ri = fmt.Sprint(rand.Float64() * float64(param))
					break

				default:
					fmt.Println("Couldn't resolve expression :", expr)
					fmt.Println("Do you wish to ignore ? y/n")
					var ans string
					_, e := fmt.Scan(&ans)
					exceptionChecker(e)
					ans = strings.ToLower(ans)
					if ans == "y" {
						continue
					} else {
						panic("Randify : Stopped by user")
					}
				} // switch args[0]

				break // case 2

			case 3: // [<indentifier> <arg1> <arg2> <arg3>]
				switch strings.ToLower(args[0]) {

				case "int":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Intn(max-min) + min)
					break

				case "int31":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Int31n(int32(max-min)) + int32(min))
					break

				case "int63":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Int63n(int64(max-min)) + int64(min))
					break

				case "uint":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Uint32()%uint32(max-min) + uint32(min))
					break

				case "uint32":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Uint32()%uint32(max-min) + uint32(min))
					break

				case "uint64":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Uint32()%uint32(max-min) + uint32(min))
					break

				case "float":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Float32()*float32(max-min) + float32(min))
					break

				case "float32":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Float32()*float32(max-min) + float32(min))
					break

				case "float64":
					min, err := strconv.Atoi(args[1])
					exceptionChecker(err)
					max, err2 := strconv.Atoi(args[2])
					exceptionChecker(err2)
					positiveChecker(expr, max-min)
					ri = fmt.Sprint(rand.Float64()*float64(max-min) + float64(min))
					break

				default:
					fmt.Println("Couldn't resolve expression :", expr)
					fmt.Println("Do you wish to ignore ? y/n")
					var ans string
					_, e := fmt.Scan(&ans)
					exceptionChecker(e)
					ans = strings.ToLower(ans)
					if ans == "y" {
						continue
					} else {
						panic("Randify : Stopped by user")
					}

				}
				break

			default:
				fmt.Println("unexpected expression length in :", expr)
				fmt.Println("Do you wish to ignore ? y/n")
				var ans string
				_, e := fmt.Scan(&ans)
				exceptionChecker(e)
				ans = strings.ToLower(ans)
				if ans == "y" {
					continue
				} else {
					panic("Randify : Stopped by user")
				}
			} // switch len(args)
			replacedData = strings.Replace(replacedData, expr, ri, 1)
		}
	}
	return replacedData
}
