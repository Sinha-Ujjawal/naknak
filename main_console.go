package main

import (
	"fmt"
	"naknak/naknak"
	"os"
)

func usage(program string) {
	fmt.Printf("Usage: %s <SUBCOMMAND> <PARAMS...>\n", program)
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("  encode <TEXT>: encode the given text into naknak language")
	fmt.Println("  decode <TEXT>: decode the naknak text")
	os.Exit(1)
}

func main() {
	if len(os.Args) == 0 {
		fmt.Printf("ERROR: this is unexpected, os.Args is empty!\n")
		os.Exit(1)
	}
	program := os.Args[0]
	os.Args = os.Args[1:]
	if len(os.Args) == 0 {
		fmt.Println("ERROR: No commands provided!")
		usage(program)
	}
	subcommand := os.Args[0]
	os.Args = os.Args[1:]
	switch subcommand {
	case "encode":
		if len(os.Args) == 0 {
			fmt.Println("ERROR: text not provided!")
			usage(program)
		}
		text := os.Args[0]
		encodedText := naknak.Encode(text)
		fmt.Println(encodedText)
	case "decode":
		if len(os.Args) == 0 {
			fmt.Println("ERROR: text not provided!")
			usage(program)
		}
		encodedText := os.Args[0]
		decodedText := naknak.Decode(encodedText)
		fmt.Println(decodedText)
	default:
		fmt.Printf("ERROR: Unknown subcommand: `%s`!\n", subcommand)
		usage(program)
	}
}
