package main

import (
	"fmt"
	"os"
	"go-reloaded/hex"
	"go-reloaded/low"
	"go-reloaded/up"
	"go-reloaded/bin"
	"go-reloaded/caps"
	"go-reloaded/punctuation"
	"go-reloaded/vowel"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input-file> <output-file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input text from file
	inputContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Apply modifications based on command-line arguments
	modifiedText := string(inputContent)
	modifiedText = hex.ApplyHexModifications(modifiedText)
	modifiedText = bin.ApplyBinModifications(modifiedText)
	modifiedText = up.ApplyUpModifications(modifiedText)
	modifiedText = low.ApplyLowModifications(modifiedText)
	modifiedText = caps.ApplyCapModifications(modifiedText)
	modifiedText = vowel.ApplyVowelModifications(modifiedText)
	modifiedText = punctuation.FormatPunctuation(modifiedText)

	// Write modified text to output file
	err = os.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Modifications applied successfully. Output written to", outputFile)
}
