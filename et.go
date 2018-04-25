package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	minwidth := flag.Int("m", 0, "minimum width")
	tabwidth := flag.Int("t", 0, "tab width")
	padding := flag.Int("p", 3, "padding width")

	flag.Parse()

	w := tabwriter.NewWriter(os.Stdout, *minwidth, *tabwidth, *padding, ' ', 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}

	w.Flush()

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
