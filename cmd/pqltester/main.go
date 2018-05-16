package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	yml := flag.String("yml", "", ".yml file with tests")
	tests := flag.String("tests", "", ".pql test file")

	flag.Parse()

	if *yml == "" {
		log.Fatalf("no .yml file provided.")
	}

	if *tests == "" {
		log.Fatalf("no .pql file provided.")
	}

	os.Exit(0)
}
