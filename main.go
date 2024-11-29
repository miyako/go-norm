package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/unicode/norm"
)

func main() {

	format := flag.String("format", "NFD", "format")

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	str := string(stdin)

	switch *format {
	case "NFC":
		NFC := norm.NFC.String(str)
		fmt.Print(NFC)
	case "NFKC":
		NFKC := norm.NFKC.String(str)
		fmt.Print(NFKC)
	case "NFKD":
		NFKD := norm.NFKD.String(str)
		fmt.Print(NFKD)
	default:
		NFD := norm.NFD.String(str)
		fmt.Print(NFD)
	}

}
