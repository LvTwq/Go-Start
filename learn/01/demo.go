package main

import (
	"flag"
	"fmt"
	"myGoProject/learn/02/lib"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "lvmc", "The greeting object.")

	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Println(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
	lib.Hello(name)
}
