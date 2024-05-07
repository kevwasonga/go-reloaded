package main

import (
    "testing"
    "go-reloaded/caps"
    "go-reloaded/up"
    "go-reloaded/low"
    "go-reloaded/hex"
    "go-reloaded/bin"
    "go-reloaded/vowel"
    "go-reloaded/punctuation"
)

func TestApplyCapModifications(t *testing.T) {
    // Define test cases for caps.ApplyCapModifications
    capTestCases := []struct {
        input    string
        expected string
    }{
        {"this is so exciting (cap, 2)", "this is So Exciting"},
        {"this is so exciting (cap)", "this is so Exciting"},
        // Add more test cases as needed
    }

    // Iterate through test cases for caps.ApplyCapModifications
    for _, tc := range capTestCases {
        // Call the function being tested
        result := caps.ApplyCapModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyCapModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyUpModifications(t *testing.T) {
    // Define test cases for up.ApplyUpModifications
    upTestCases := []struct {
        input    string
        expected string
    }{
        {"this is so exciting (up, 2)", "this is SO EXCITING"},
        {"this is so exciting (up)", "this is so EXCITING"},
        // Add more test cases as needed
    }

    // Iterate through test cases for up.ApplyUpModifications
    for _, tc := range upTestCases {
        // Call the function being tested
        result := up.ApplyUpModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyUpModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyLowModifications(t *testing.T) {
    // Define test cases for low.ApplyLowModifications
    lowTestCases := []struct {
        input    string
        expected string
    }{
        {"THIS IS SO EXCITING (low, 2)", "THIS IS so exciting"},
        {"THIS IS SO EXCITING (low)", "THIS IS SO exciting"},
        // Add more test cases as needed
    }

    // Iterate through test cases for low.ApplyLowModifications
    for _, tc := range lowTestCases {
        // Call the function being tested
        result := low.ApplyLowModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyLowModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyHexModifications(t *testing.T) {
    // Define test cases for hex.ApplyHexModifications
    hexTestCases := []struct {
        input    string
        expected string
    }{
        {"1E (hex) files were added", "30 files were added"},
        // Add more test cases as needed
    }

    // Iterate through test cases for hex.ApplyHexModifications
    for _, tc := range hexTestCases {
        // Call the function being tested
        result := hex.ApplyHexModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyHexModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyBinModifications(t *testing.T) {
    // Define test cases for bin.ApplyBinModifications
    binTestCases := []struct {
        input    string
        expected string
    }{
        {"It has been 10 (bin) years", "It has been 2 years"},
        // Add more test cases as needed
    }

    // Iterate through test cases for bin.ApplyBinModifications
    for _, tc := range binTestCases {
        // Call the function being tested
        result := bin.ApplyBinModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyBinModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyVowelModifications(t *testing.T) {
    // Define test cases for vowel.ApplyVowelModifications
    vowelTestCases := []struct {
        input    string
        expected string
    }{
        {"There it was. A amazing rock!", "There it was. An amazing rock!"},
        {"There it was. a amazing rock!", "There it was. an amazing rock!"},
        // Add more test cases as needed
    }

    // Iterate through test cases for vowel.ApplyVowelModifications
    for _, tc := range vowelTestCases {
        // Call the function being tested
        result := vowel.ApplyVowelModifications(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("ApplyVowelModifications(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}

func TestApplyPunctuationModifications(t *testing.T) {
    // Define test cases for punctuation.ApplyPunctuationModifications
    punctuationTestCases := []struct {
        input    string
        expected string
    }{
        {"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
        {"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
        // Add more test cases as needed
    }

    // Iterate through test cases for punctuation.ApplyPunctuationModifications
    for _, tc := range punctuationTestCases {
        // Call the function being tested
        result := punctuation.FormatPunctuation(tc.input)

        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("FormatPunctuation(%q) = %q; want %q", tc.input, result, tc.expected)
        }
    }
}
