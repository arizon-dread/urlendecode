package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	enc := flag.String("e", "", "pass the string to encode")
	dec := flag.String("d", "", "pass the string to decode")
	flag.Parse()
	if *enc != "" {
		fmt.Printf("%v\n", url.QueryEscape(*enc))
		os.Exit(0)
	} else if *dec != "" {
		res, err := url.QueryUnescape(*dec)
		if err == nil {
			fmt.Printf("%v\n", res)
			os.Exit(0)
		}
	}
	flag.PrintDefaults()
	os.Exit(1)
}
