package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "number":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		var i, _ = strconv.Atoi(flag.Arg(1))
		var result string = fizzbuzz(i)

		fmt.Println(result)
	}
}

func fizzbuzz(number int) (result string) {
	switch {
	case number == 0:
		return "0"
	case number%15 == 0:
		return "fizzbuzz"
	case number%3 == 0:
		return "fizz"
	case number%5 == 0:
		return "buzz"
	}

	return strconv.Itoa(number)
}
