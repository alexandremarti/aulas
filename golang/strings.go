package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	textPtr := flag.String("texto", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	nroPtr := flag.Float64("nro", 0, "número")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t, nro: %.2f\n", *textPtr, *metricPtr, *uniquePtr, *nroPtr)
}
