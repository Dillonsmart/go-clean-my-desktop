package main

import (
	"flag"
)

func main() {
	stepThrough := flag.Bool("step", false, "Step through the files")

	flag.Parse()

	Run(*stepThrough)
}
