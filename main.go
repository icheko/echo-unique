package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const currentVersion string = "1.0"
const packageName string = "echo-unique"

var (
	help    = flag.Bool("help", false, "Show help screen.")
	version = flag.Bool("version", false, "Display version info.")
)

func main() {
	flag.Parse()

	if *help {
		printUsage()
		return
	}

	if *version {
		fmt.Println(packageName + " version " + currentVersion)
		return
	}

	uniqueStrings := make(map[string]struct{})

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cat file | " + packageName)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = removeNewLineChar(input)
		_, found := uniqueStrings[input]

		if !found {
			fmt.Println(input)
			uniqueStrings[input] = struct{}{}
		}
	}
}

func printUsage() {
	fmt.Print(packageName + " [-flag value]\n\n")
	flag.Usage()
	fmt.Println()
}

func removeNewLineChar(str string) string {
	if strings.Contains(str, "\n") {
		return strings.Replace(str, "\n", "", 1)
	}

	return str
}
