This is a Go program that provides various text formatting functionalities such as modifying case, handling punctuation, and converting hexadecimal and binary representations.
Features

    Case Modification: Convert text to lowercase, uppercase, or capitalize specific words based on user-defined rules.
    Punctuation Formatting: Ensure proper spacing and placement of punctuation marks according to specified rules.
    Hexadecimal Conversion: Convert hexadecimal numbers to decimal representation.
    Binary Conversion: Convert binary numbers to decimal representation.

Installation

To use this tool, you need to have Go installed on your system. You can then install the tool using go get:

bash

go get github.com/username/go-reloaded

Usage

    Import the necessary packages in your Go program:

    go


Use the provided functions to format your text:

go

    // Example usage of hex package
    modifiedText := hex.ApplyHexModifications("1E (hex) files were added")

    // Example usage of punctuation package
    modifiedText := punctuation.FormatPunctuation("I was sitting over there ,and then BAMM !!")

    // Example usage of up package
    modifiedText := up.ApplyUpModifications("This is so exciting (up, 2)")

    // Example usage of low package
    modifiedText := low.ApplyLowModifications("This is SO EXCITING (low, 2)")

    // Example usage of bin package
    modifiedText := bin.ApplyBinModifications("101 (bin) files were uploaded")

    // Example usage of caps package
    modifiedText := caps.ApplyCapModifications("this is so cool (cap, 3)")

